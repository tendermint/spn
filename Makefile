#! /usr/bin/make -f

PACKAGES=$(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags 2> /dev/null || echo "dev-$(shell git describe --always)") | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
DOCKER := $(shell which docker)
COVER_FILE := coverage.txt
COVER_HTML_FILE := cover.html
BINDIR ?= $(GOPATH)/bin
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=spnd \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=spnd \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/cosmos/cosmos-sdk/types.reDnmString=[a-zA-Z][a-zA-Z0-9/:]{2,127}

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

all: tools install lint

###############################################################################
###                               Development                               ###
###############################################################################

## govet: Run go vet.
govet:
	@echo Running go vet...
	@go vet ./...

## govulncheck: Run govulncheck
govulncheck:
	@echo Running govulncheck...
	@go run golang.org/x/vuln/cmd/govulncheck ./...

FIND_ARGS := -name '*.go' -type f -not -name '*.pb.go' -not -name '*.pb.gw.go'

## format: Run gofumpt and goimports.
format:
	@echo Formatting...
	@go install mvdan.cc/gofumpt
	@go install golang.org/x/tools/cmd/goimports
	@find . $(FIND_ARGS) | xargs gofumpt -w .
	@find . $(FIND_ARGS) | xargs goimports -w -local github.com/ignite/modules

## lint: Run Golang CI Lint.
lint:
	@echo Running gocilint...
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@golangci-lint run --out-format=tab --issues-exit-code=0

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)", or just run 'make' for install"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: lint format govet govulncheck help

###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/spnd.exe ./cmd/spnd
else
	go build $(BUILD_FLAGS) -o build/spnd ./cmd/spnd
endif

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install: go.sum
	go install $(BUILD_FLAGS) ./cmd/spnd

build-reproducible: go.sum
	$(DOCKER) rm latest-build || true
	$(DOCKER) run --volume=$(CURDIR):/sources:ro \
        --env TARGET_PLATFORMS='linux/amd64 darwin/amd64 linux/arm64' \
        --env APP=spnd \
        --env VERSION=$(VERSION) \
        --env COMMIT=$(COMMIT) \
        --name latest-build cosmossdk/rbuilder:latest
	$(DOCKER) cp -a latest-build:/home/builder/artifacts/ $(CURDIR)/

###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	@go mod tidy

clean:
	rm -rf build/

.PHONY: go-mod-cache clean

###############################################################################
###                                  Test                                   ###
###############################################################################

## test-unit: Run the unit tests.
test-unit:
	@echo Running unit tests...
	@VERSION=$(VERSION) go test -mod=readonly -v -timeout 30m $(PACKAGES)

## test-race: Run the unit tests checking for race conditions
test-race:
	@echo Running unit tests with race condition reporting...
	@VERSION=$(VERSION) go test -mod=readonly -v -race -timeout 30m  $(PACKAGES)

## test-cover: Run the unit tests and create a coverage html report
test-cover:
	@echo Running unit tests and creating coverage report...
	@VERSION=$(VERSION) go test -mod=readonly -v -timeout 30m -coverprofile=$(COVER_FILE) -covermode=atomic $(PACKAGES)
	@go tool cover -html=$(COVER_FILE) -o $(COVER_HTML_FILE)
	@rm $(COVER_FILE)

## bench: Run the unit tests with benchmarking enabled
bench:
	@echo Running unit tests with benchmarking...
	@VERSION=$(VERSION) go test -mod=readonly -v -timeout 30m -bench=. $(PACKAGES)

## test: Run unit and integration tests.
test: govet test-unit

.PHONY: test test-unit test-race test-cover bench

###############################################################################
###                                Protobuf                                 ###
###############################################################################

proto-all: proto-format proto-lint proto-gen-gogo

proto-gen-gogo:
	@echo "Generating Protobuf Files"
	@buf generate --template $(CURDIR)/proto/buf.gen.gogo.yaml --output $(CURDIR)/gen/go
	@cp -r gen/go/github.com/ignite/modules/x ./
	@rm -R gen/go

proto-gen-swagger:
	@echo "Generating Protobuf Swagger"
	@buf generate --template $(CURDIR)/proto/buf.gen.swagger.yaml --output $(CURDIR)/gen/swagger

proto-gen-ts:
	@echo "Generating Protobuf Typescript"
	@buf generate --template $(CURDIR)/proto/buf.gen.ts.yaml --output $(CURDIR)/gen/ts

proto-format:
	@echo "Formatting Protobuf Files"
	@buf format --write

proto-lint:
	@echo "Linting Protobuf Files"
	@buf lint

.PHONY: proto-all proto-gen-gogo proto-gen-swagger proto-gen-ts proto-format proto-lint

###############################################################################
###                               Simulation                                ###
###############################################################################

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

SIMAPP = ./app
SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 100
SIM_CI_NUM_BLOCKS ?= 200
SIM_CI_BLOCK_SIZE ?= 26
SIM_PERIOD ?= 50
SIM_COMMIT ?= true
SIM_TIMEOUT ?= 24h

# test-sim-nondeterminism: Run simulation test checking for app state nondeterminism
test-sim-nondeterminism:
	@echo "Running non-determinism test..."
	@VERSION=$(VERSION) go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -Period=$(SIM_PERIOD)  \
		-v -timeout $(SIM_TIMEOUT)

# test-sim-import-export: Run simulation test checking import and export app state determinism
# go get github.com/cosmos/tools/cmd/runsim@v1.0.0
test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 2 2 TestAppImportExport

# test-sim-after-import: Run simulation test checking import after simulation
# go get github.com/cosmos/tools/cmd/runsim@v1.0.0
test-sim-after-import: runsim
	@echo "Running application simulation-after-import. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 2 2 TestAppSimulationAfterImport

# test-sim-nondeterminism-long: Run simulation test checking for app state nondeterminism with a big data
test-sim-nondeterminism-long:
	@echo "Running non-determinism test..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=100 -Commit=true -Period=0 -v -timeout 1h

# test-sim-import-export-long: Test import/export simulation
test-sim-import-export-long: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 5 5 TestAppImportExport

# test-sim-import-export-long: Test import/export simulation with a big data
test-sim-after-import-long: runsim
	@echo "Running application simulation-after-import. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 5 5 TestAppSimulationAfterImport

# test-sim-ci: Run lightweight simulation for CI pipeline
test-sim-ci:
	@echo "Running application benchmark for numBlocks=$(SIM_CI_NUM_BLOCKS), blockSize=$(SIM_CI_BLOCK_SIZE)"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_CI_NUM_BLOCKS) -BlockSize=$(SIM_CI_BLOCK_SIZE) -Commit=$(SIM_COMMIT) \
		-Period=$(SIM_PERIOD) -timeout $(SIM_TIMEOUT)

# test-sim-benchmark: Run heavy benchmarking simulation
test-sim-benchmark:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Period=$(SIM_PERIOD) \
		-Commit=$(SIM_COMMIT) timeout $(SIM_TIMEOUT)

# test-sim-benchmark: Run heavy benchmarking simulation with CPU and memory profiling
test-sim-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$ \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Period=$(SIM_PERIOD) \
		-Commit=$(SIM_COMMIT) timeout $(SIM_TIMEOUT)-cpuprofile cpu.out -memprofile mem.out

.PHONY: \
test-sim-nondeterminism \
test-sim-nondeterminism-long \
test-sim-import-export \
test-sim-import-export-long \
test-sim-after-import \
test-sim-after-import-long \
test-sim-ci \
test-sim-profile \
test-sim-benchmark

.DEFAULT_GOAL := install

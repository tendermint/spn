#! /usr/bin/make -f

PACKAGES=$(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags 2> /dev/null || echo "dev-$(shell git describe --always)") | sed 's/^v//')
SIMAPP = ./app
DOCKER := $(shell which docker)
COVER_FILE := coverage.txt
COVER_HTML_FILE := cover.html

## govet: Run go vet.
govet:
	@echo Running go vet...
	@go vet ./...

FIND_ARGS := -name '*.go' -type f -not -name '*.pb.go' -not -name '*.pb.gw.go'

## format: Run gofmt and goimports.
format:
	@echo Formatting...
	@go install mvdan.cc/gofumpt
	@go install golang.org/x/tools/cmd/goimports
	@find . $(FIND_ARGS) | xargs gofumpt -w .
	@find . $(FIND_ARGS) | xargs goimports -w -local github.com/tendermint/spn

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

.PHONY: lint format govet help

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

proto-all: proto-format proto-gen

protoVer=v0.7
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen=ignite-spn-proto-gen-$(protoVer)
containerProtoFmt=ignite-spn-proto-fmt-$(protoVer)

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protocgen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi


SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 100
SIM_CI_NUM_BLOCKS ?= 200
SIM_CI_BLOCK_SIZE ?= 26
SIM_PERIOD ?= 50
SIM_COMMIT ?= true
SIM_TIMEOUT ?= 24h

## test-sim-nondeterminism: Run simulation test checking for app state nondeterminism
test-sim-nondeterminism:
	@echo "Running non-determinism test..."
	@VERSION=$(VERSION) go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -Period=$(SIM_PERIOD)  \
		-v -timeout $(SIM_TIMEOUT)

## test-sim-ci: Run lightweight simulation for CI pipeline
test-sim-ci:
	@echo "Running application benchmark for numBlocks=$(SIM_CI_NUM_BLOCKS), blockSize=$(SIM_CI_BLOCK_SIZE)"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_CI_NUM_BLOCKS) -BlockSize=$(SIM_CI_BLOCK_SIZE) -Commit=$(SIM_COMMIT) \
		-Period=$(SIM_PERIOD) -timeout $(SIM_TIMEOUT)

## test-sim-benchmark: Run heavy benchmarking simulation
test-sim-benchmark:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Period=$(SIM_PERIOD) \
		-Commit=$(SIM_COMMIT) timeout $(SIM_TIMEOUT)

## test-sim-benchmark: Run heavy benchmarking simulation with CPU and memory profiling
test-sim-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$ \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Period=$(SIM_PERIOD) \
		-Commit=$(SIM_COMMIT) timeout $(SIM_TIMEOUT)-cpuprofile cpu.out -memprofile mem.out

.PHONY: test-sim-nondeterminism test-sim-ci test-sim-profile test-sim-benchmark

.DEFAULT_GOAL := install

#! /usr/bin/make -f

PACKAGES=$(shell go list ./...)
VERSION := $(shell echo $(shell git describe --tags 2> /dev/null || echo "dev-$(shell git describe --always)") | sed 's/^v//')
SIMAPP = ./app
COVER_FILE := coverage.txt
COVER_HTML_FILE := cover.html

## govet: Run go vet.
govet:
	@echo Running go vet...
	@go vet ./...

FIND_ARGS := -name '*.go' -type f -not -name '*.pb.go'

## format: Run gofmt and goimports.
format:
	@echo Formatting...
	@find . $(FIND_ARGS) | xargs gofmt -d -s
	@find . $(FIND_ARGS) | xargs goimports -w -local github.com/tendermint/spn

## lint: Run Golang CI Lint.
lint:
	@echo Running gocilint...
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

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_PERIOD ?= 10
SIM_COMMIT ?= true

## test-sim-nondeterminism: Run simulation test checking for app state nondeterminism
test-sim-nondeterminism:
	@echo "Running non-determinism test..."
	@VERSION=$(VERSION) go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=10 -BlockSize=100 -Commit=true -Period=0 -v -timeout 24h

## test-sim-ci: Run lightweight simulation for CI pipeline
test-sim-ci:
	@echo "Running application benchmark for numBlocks=200, blockSize=26"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$  \
		-Enabled=true -NumBlocks=200 -BlockSize=26 -Commit=$(SIM_COMMIT) -timeout 24h

## test-sim-benchmark: Run heavy benchmarking simulation
test-sim-benchmark:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Period=$(SIM_PERIOD) \
		-Commit=$(SIM_COMMIT) -timeout 24h

## test-sim-benchmark: Run heavy benchmarking simulation with CPU and memory profiling
test-sim-profile:
	@echo "Running application benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@VERSION=$(VERSION) go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkSimulation$$ \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Period=$(SIM_PERIOD) \
		-Commit=$(SIM_COMMIT) -timeout 24h -cpuprofile cpu.out -memprofile mem.out

.PHONY: test-sim-nondeterminism test-sim-ci test-sim-profile test-sim-benchmark

.DEFAULT_GOAL := install

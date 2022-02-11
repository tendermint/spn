#! /usr/bin/make -f

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

## test-unit: Run the unit tests.
test-unit:
	@echo Running unit tests...
	@go test -v ./...

## test: Run unit and integration tests.
test: govet test-unit

help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)", or just run 'make' for install"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.DEFAULT_GOAL := install

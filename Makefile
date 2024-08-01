
.PHONY: build

VERSION=$(shell git describe --tags)
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_COMMIT_DATE=$(shell git log -n1 --pretty='format:%cd' --date=format:'%Y%m%d')
REPO=github.com/zkMeLabs/mechain-scan-backend
IMAGE_NAME=ghcr.io/bnb-chain/mechain-scanner

ldflags = -X $(REPO)/version.AppVersion=$(VERSION) \
          -X $(REPO)/version.GitCommit=$(GIT_COMMIT) \
          -X $(REPO)/version.GitCommitDate=$(GIT_COMMIT_DATE)

build:
	go build -o build/scand cmd/scand/*.go
 
###############################################################################
###                                Linting                                  ###
###############################################################################

golangci_lint_cmd=golangci-lint
golangci_version=v1.59.1

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --timeout=10m

lint-fix:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

format:
	bash scripts/format.sh

.PHONY: lint lint-fix format

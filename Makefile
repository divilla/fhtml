GODIRS=$(shell go list -f {{.Dir}} ./...)

.PHONY: race
race: ## Run data race detector
	@go test -race -short $(shell go list ./... | grep -v /vendor/)

.PHONY: msan
msan: ## Run memory sanitizer
	@go env -w "CC=clang"
	@go test -msan -short $(shell go list ./... | grep -v /vendor/)

.PHONY: lint-upgrade
lint-upgrade:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.62.0
	@golangci-lint --version

.PHONY: lint
lint:  ## lint
	@golangci-lint run --timeout 5m --print-issued-lines=false --out-format code-climate:gl-code-quality-report.json,line-number

.PHONY: fmt
fmt:  ## gofmt & goimports
	@go mod tidy
	@goimports -l -w $(GODIRS)

.PHONY: benchmark
bench:  ## go test
	@go test -bench=. -benchmem cmd/*.go

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


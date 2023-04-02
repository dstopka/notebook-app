BASE_DIR=$(shell git rev-parse --show-toplevel)

.PHONY: fmt
fmt:
	@gofmt -s -w .

.PHONY: fmt-check
fmt-check:
	@$(BASE_DIR)/scripts/fmt-check.sh

.PHONY: test
test:
	@go test -v ./...

.PHONY: unit-test
unit-test:
	@go test -v -short ./...

.PHONY: integration-test
integration-test:
	@go test -v -run Integration ./...

.PHONY: component-test
component-test:
	@go test -v -run Component ./...

.PHONY: proto
proto:
	@$(BASE_DIR)/scripts/genproto.sh notebooks
	
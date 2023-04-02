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
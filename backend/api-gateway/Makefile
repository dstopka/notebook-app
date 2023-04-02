BASE_DIR=$(shell git rev-parse --show-toplevel)

.PHONY: fmt
fmt:
	@gofmt -s -w .

.PHONY: test
test:
	@go test -v ./...

.PHONY: openapi
openapi:
	@go generate -run="oapi-codegen"
	
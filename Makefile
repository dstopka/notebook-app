BASE_DIR=$(shell pwd)

.PHONY: all
all: 

.PHONY: test
test:
	@echo "Running tests for backend/common"
	@make -C $(BASE_DIR)/backend/common test

	@echo "Running tests for backend/api-gateway"
	@make -C $(BASE_DIR)/backend/api-gateway test

	@echo "Running tests for backend/notebooks"
	@make -C $(BASE_DIR)/backend/notebooks test

	@echo "Running tests for backend/users"
	@make -C $(BASE_DIR)/backend/users test

.PHONY: fmt
fmt:
	@echo "Formatting *.go files in backend/common"
	@make -C $(BASE_DIR)/backend/common fmt

	@echo "Formatting *.go files in backend/api-gateway"
	@make -C $(BASE_DIR)/backend/api-gateway fmt

	@echo "Formatting *.go files in backend/notebooks"
	@make -C $(BASE_DIR)/backend/notebooks fmt

	@echo "Formatting *.go files in backend/users"
	@make -C $(BASE_DIR)/backend/users fmt

.PHONY: fmt-check
fmt-check:
	@echo "Checking format of *.go files in backend/common"
	@make -C $(BASE_DIR)/backend/common fmt-check

	@echo "Checking format of *.go files in backend/api-gateway"
	@make -C $(BASE_DIR)/backend/api-gateway fmt-check

	@echo "Checking format of *.go files in backend/notebooks"
	@make -C $(BASE_DIR)/backend/notebooks fmt-check
	
	@echo "Checking format of *.go files in backend/users"
	@make -C $(BASE_DIR)/backend/users fmt-check

.PHONY: proto
proto:
	@make -C $(BASE_DIR)/backend/notebooks proto
	@make -C $(BASE_DIR)/backend/users proto

.PHONY: openapi
openapi:
	@make -C $(BASE_DIR)/backend/api-gateway openapi
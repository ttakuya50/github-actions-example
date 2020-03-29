MAKEFILE_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

lint:
	docker run --rm -v $(MAKEFILE_DIR):/app -w /app golangci/golangci-lint:v1.23.8 golangci-lint run --config=.github/.golangci.yml
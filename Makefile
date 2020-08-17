.PHONY: deps
deps:
	@go mod download
	@go mod vendor
	@go mod tidy

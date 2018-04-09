.PHONY: default
default: deps lint migrate tests install

.PHONY: install
install:
	@go install ./...

.PHONY: tests
tests:
	@go test -v ./...
.PHONY: default
default: deps install

.PHONY: deps
deps:
	@go get github.com/golang/dep/cmd/dep
	@dep ensure

.PHONY: install
install:
	@go install ./...
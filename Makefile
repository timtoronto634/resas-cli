.PHONY: lint tools.get
GOBIN = $(shell go env GOPATH)/bin

tools.get:
	go install github.com/mgechev/revive@latest

lint:
	$(GOBIN)/revive

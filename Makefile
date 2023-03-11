.PHONY: lint tools.get vet
GOBIN = $(shell go env GOPATH)/bin

tools.get:
	go install github.com/mgechev/revive@latest

revive:
	$(GOBIN)/revive -config revive.toml -formatter friendly ./...

vet:
	go vet ./...

lint: vet revive
	

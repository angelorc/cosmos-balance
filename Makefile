GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

swagger-ui:
	@echo "generate cosmos tracker swagger..."
	@swag init -g ./server/server.go --output ./swagger

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building ctracker binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/ctracker.exe main.go
else
	@echo "building ctracker binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/ctracker main.go
endif

install: go.sum
	@echo "installing ctracker binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o $(GOBIN)/ctracker main.go

.PHONY: build install swagger-ui
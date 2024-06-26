.PHONY: gen-proto build testacc

BIN_OUT_DIR := out
BINARY := $(BIN_OUT_DIR)/terraform-provider-bria
HOSTNAME=galoymoney

PROTO_DIR := proto/vendor
PROTO_OUTPUT_DIR := bria/proto

version = 0.0.14
os_arch = $(shell go env GOOS)_$(shell go env GOARCH)
provider_path = registry.terraform.io/galoymoney/bria/$(version)/$(os_arch)/

fmt:
	goimports -l -w .
	go mod tidy
	terraform fmt --recursive
	nix fmt

gen-proto:
	mkdir -p $(PROTO_OUTPUT_DIR)
	protoc -I $(PROTO_DIR) \
		--go_out=$(PROTO_OUTPUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUTPUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/api/bria.proto

gen-docs:
	 go generate ./...

build:
	go build -o $(BINARY) main.go

install: gen-proto build
	mkdir -p ~/.terraform.d/plugins/${provider_path}
	mv ${BINARY} ~/.terraform.d/plugins/${provider_path}
	rm -rf example/.terraform example/.terraform.lock.hcl example/terraform.tfstate*

clean:
	rm -rf ~/.terraform.d/plugins/${provider_path}

testacc:
	TF_ACC=1 go test -v ./...

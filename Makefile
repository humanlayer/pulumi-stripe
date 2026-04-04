PROJECT_NAME := stripe
PROVIDER := pulumi-resource-$(PROJECT_NAME)
TFGEN := pulumi-tfgen-$(PROJECT_NAME)
VERSION ?= 0.0.1-dev

GOPATH := $(shell go env GOPATH)
PROVIDER_PATH := provider

.PHONY: tfgen provider build_nodejs clean

# Generate schema.json and bridge-metadata.json
tfgen:
	cd $(PROVIDER_PATH) && go build -o ../bin/$(TFGEN) ./cmd/$(TFGEN)
	./bin/$(TFGEN) schema --out provider/cmd/$(PROVIDER)

# Build the provider binary
provider:
	cd $(PROVIDER_PATH) && go build -ldflags "-X github.com/humanlayer/pulumi-stripe/provider/pkg/version.Version=$(VERSION)" -o ../bin/$(PROVIDER) ./cmd/$(PROVIDER)

# Generate TypeScript SDK
build_nodejs:
	./bin/$(TFGEN) nodejs --out sdk/nodejs

# Build everything
build: tfgen provider build_nodejs
	cd sdk/nodejs && npm install && npm run build

# Clean generated files
clean:
	rm -rf bin/ sdk/nodejs/

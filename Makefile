PACK             := terraform
ORG              := kmochejski
PROJECT          := github.com/${ORG}/pulumi-${PACK}
NODE_MODULE_NAME := @pulumi/${PACK}
TF_NAME          := ${PACK}
PROVIDER_PATH    := provider/v5
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell pulumictl get version)

WORKING_DIR     := $(shell pwd)

.PHONY: development provider build_sdks build_nodejs cleanup

development:: provider build_sdks install_sdks cleanup # Build the provider & SDKs for a development environment

# Required for the codegen action that runs in pulumi/pulumi and pulumi/pulumi-terraform-bridge
build:: provider build_sdks install_sdks
only_build:: build

provider:: # build the provider binary
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: provider build_nodejs build_go

build_nodejs:: VERSION := $(shell pulumictl get version --language javascript)
build_nodejs:: # build the node sdk
	cd sdk/nodejs/ && \
        yarn install && \
        yarn run tsc && \
        cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
    	sed -i.bak -e "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json

build_go:: # build the go sdk

lint_provider:: provider # lint the provider code
	cd provider && golangci-lint run -c ../.golangci.yml

cleanup:: # cleans up the temporary directory
	rm -r $(WORKING_DIR)/bin

help::
	@grep '^[^.#]\+:\s\+.*#' Makefile | \
 	sed "s/\(.\+\):\s*\(.*\) #\s*\(.*\)/`printf "\033[93m"`\1`printf "\033[0m"`	\3 [\2]/" | \
 	expand -t20

clean::

install_nodejs_sdk::
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

install_go_sdk::

install_sdks:: install_nodejs_sdk

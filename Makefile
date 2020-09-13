#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
#SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
BUILDDIR ?= $(CURDIR)/build


export GO111MODULE = on

# process build tags

build_tags = netgo
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags
ldflags := -X "github.com/cosmos/cosmos-sdk/version.Name=oin" \
   -X "github.com/cosmos/cosmos-sdk/version.ServerName=oind" \
   -X "github.com/cosmos/cosmos-sdk/version.ClientName=oincli" \
   -X "github.com/cosmos/cosmos-sdk/version.Version=$(VERSION)" \
   -X "github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)" \
   -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
all: install

build: go.sum
	mkdir -p $(BUILDDIR)
	go build  $(BUILD_FLAGS) -o $(BUILDDIR)/ ./cmd/oind
	go build  $(BUILD_FLAGS) -o $(BUILDDIR)/ ./cmd/oincli

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install: go.sum
		go install  $(BUILD_FLAGS) ./cmd/oind
		go install  $(BUILD_FLAGS) ./cmd/oincli
go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		@go mod verify

clean:
	@rm -rf build/

distclean: clean
	rm -rf vendor/

###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs goimports -w -local github.com/cosmos/gaia

.PHONY: all build-linux install format lint \
	go-mod-cache clean build \


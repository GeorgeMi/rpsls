
# Version Vars
VERSION_TAG := $(shell git describe --tags --always)
VERSION_VERSION := $(shell git log --date=iso --pretty=format:"%cd" -1) $(VERSION_TAG)
VERSION_COMPILE := $(shell date +"%F %T %z") by $(shell go version)
VERSION_BRANCH  := $(shell git rev-parse --abbrev-ref HEAD)
VERSION_GIT_DIRTY := $(shell git diff --no-ext-diff 2>/dev/null | wc -l | awk '{print $1}')
VERSION_DEV_PATH:= $(shell pwd)

# Go Checkup
GOPATH ?= $(shell go env GOPATH)
GO111MODULE:=auto
export GO111MODULE
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
PATH := ${GOPATH}/bin:$(PATH)

GO = go

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m➡\033[0m")

# Commands
.PHONY: all

.PHONY: build
build: ; $(info $(M) Building executable...) @ ## Build program binary
	$Q mkdir -p bin
	$Q ret=0 && for d in $$($(GO) list -f '{{if (eq .Name "main")}}{{.ImportPath}}{{end}}' ./...); do \
		b=$$(basename $${d}) ; \
		$(GO) mod tidy ; \
		$(GO) build -o bin/$${b} $$d || ret=$$? ; \
		echo "$(M) Build: bin/$${b}" ; \
		echo "$(M) Done!" ; \
	done ; exit $$ret

.PHONY: run
run: ; $(info $(M) Running dev build (on the fly) ...) @ ## Run intermediate builds
	$Q $(GO) run -race bin/rpsls-api


help:
	$Q echo "\nGo Clean Arch\n----------------"
	$Q grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

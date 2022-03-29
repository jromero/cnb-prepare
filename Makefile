BIN_NAME=preparer

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')

.PHONY: default
default: test

.PHONY: help
help:
	@echo 'Management commands for preparer:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make test            Run tests.'
	@echo '    make clean           Clean the directory tree.'
	@echo

.PHONY: build
build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/jromero/cnb-prepare/pkg/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/jromero/cnb-prepare/pkg/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME} cmd/preparer/preparer.go

.PHONY: clean
clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

.PHONY: test
test:
	go test -v -count 1 -p 1 ./...

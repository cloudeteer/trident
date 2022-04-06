# Copyright 2018 NetApp, Inc. All Rights Reserved.


GOARCH ?= amd64
GOGC ?= ""
GOPROXY ?= https://proxy.golang.org
GO_IMAGE ?= golang:1.17
HELM_IMAGE = alpine/helm:3.6.1
GOLANGCI-LINT_VERSION ?= v1.31.0
TRIDENT_VOLUME = trident_build
TRIDENT_VOLUME_PATH = /go/src/github.com/netapp/trident
TRIDENT_CONFIG_PKG = github.com/netapp/trident/config
OPERATOR_CONFIG_PKG = github.com/netapp/trident/operator/config
TRIDENT_KUBERNETES_PKG = github.com/netapp/trident/persistent_store/crd
VERSION_FILE = github.com/netapp/trident/hack/VERSION
K8S_CODE_GENERATOR = code-generator-kubernetes-1.18.2
BUILD_CONTAINER_NAME ?= trident-build

## build flags variables
GITHASH ?= `git describe --match=NeVeRmAtCh --always --abbrev=40 --dirty || echo unknown`
BUILD_TYPE ?= custom
BUILD_TYPE_REV ?= 0
BUILD_TIME = `date`

# common variables
PORT ?= 8000
ROOT = $(shell pwd)
BIN_DIR = ${ROOT}/bin
MACOS_BIN_DIR = ${BIN_DIR}/macos
COVERAGE_DIR = ${ROOT}/coverage
BIN ?= trident_orchestrator
TARBALL_BIN ?= trident
CLI_BIN ?= tridentctl
CLI_PKG ?= github.com/netapp/trident/cli
K8S ?= ""
BUILD = build
VERSION ?= $(shell cat ${ROOT}/hack/VERSION)
VOLUME_ARG = -v ${TRIDENT_VOLUME}:/go
ifdef CREATE_BASE_IMAGE
	VOLUME_ARG =
endif

DR_LINUX = docker run \
	--name=trident-build \
	--net=host \
	-e CGO_ENABLED=0 \
	-e GOOS=linux \
	-e GOARCH=amd64 \
	-e GOGC="" \
	-e GOPROXY=https://proxy.golang.org \
	-e XDG_CACHE_HOME=/go/cache \
	-v trident_build:go \
	-v "..//hack/VERSION":"/go/src/github.com/netapp/trident" \
	-w /go/src/github.com/netapp/trident \
	golang:1.17

DR_MACOS = docker run \
	--name=${BUILD_CONTAINER_NAME} \
	--net=host \
	-e GOOS=darwin \
	-e GOARCH=$(GOARCH) \
	-e GOGC=$(GOGC) \
	-e GOPROXY=$(GOPROXY) \
	-e XDG_CACHE_HOME=/go/cache \
	${VOLUME_ARG} \
	-v "${ROOT}":"${TRIDENT_VOLUME_PATH}" \
	-w $(TRIDENT_VOLUME_PATH) \
	$(GO_IMAGE)

GO_CMD ?= go

GO_LINUX = ${DR_LINUX} go

GO_MACOS = ${DR_MACOS} ${GO_CMD}

DR_HELM = docker run --rm -v "${ROOT}":"/apps" $(HELM_IMAGE)

.PHONY: default build trident_build trident_build_all tridentctl_build dist dist_tar dist_tag test test_core test_other test_coverage_report clean fmt install vet install-lint lint-precommit lint-prepush mocks

default: dist

## version variables
TRIDENT_VERSION ?= ${VERSION}
TRIDENT_IMAGE ?= trident
ifeq ($(BUILD_TYPE),custom)
TRIDENT_VERSION := ${TRIDENT_VERSION}-custom
else ifneq ($(BUILD_TYPE),stable)
TRIDENT_VERSION := ${TRIDENT_VERSION}-${BUILD_TYPE}.${BUILD_TYPE_REV}
endif

## tag variables
TRIDENT_TAG := ${TRIDENT_IMAGE}:${TRIDENT_VERSION}
ifdef REGISTRY_ADDR
TRIDENT_TAG := ${REGISTRY_ADDR}/${TRIDENT_TAG}
endif
DIST_REGISTRY ?= netapp
TRIDENT_DIST_TAG := ${DIST_REGISTRY}/${TRIDENT_IMAGE}:${TRIDENT_VERSION}

## trident operator variable
OPERATOR_IMAGE ?= trident-operator
DEFAULT_TRIDENT_OPERATOR_REPO ?= netapp/${OPERATOR_IMAGE}
DEFAULT_TRIDENT_OPERATOR_VERSION ?= ${VERSION}
DEFAULT_TRIDENT_OPERATOR_IMAGE := ${DEFAULT_TRIDENT_OPERATOR_REPO}:${DEFAULT_TRIDENT_OPERATOR_VERSION}
OPERATOR_DIST_TAG := ${DIST_REGISTRY}/${OPERATOR_IMAGE}:${TRIDENT_VERSION}

# Go compiler flags need to be properly encapsulated with double quotes to handle spaces in values
BUILD_FLAGS = "-s -w -X \"github.com/netapp/trident/config.BuildHash=`git describe --match=NeVeRmAtCh --always --abbrev=40 --dirty || echo unknown`\" \
-X \github.com/netapp/trident/config.BuildType=custom\" -X \"github.com/netapp/trident/config.BuildTypeRev=0\" \
-X \"github.com/netapp/trident/config.BuildTime=`date`\" -X \"github.com/netapp/trident/config.BuildImage="sth"\" -X \
"github.com/netapp/trident/operator/config.BuildImage="sth"\""

## Trident build targets

trident_build:
	mkdir -p /bin
	chmod 777 /bin
	${GO_LINUX} build -ldflags $(BUILD_FLAGS) -o /go/src/github.com/netapp/trident/bin/trident_orchestrator

	docker commit "name" "sth"

	docker rm ${BUILD_CONTAINER_NAME}
	${GO_LINUX} build -ldflags $(BUILD_FLAGS) -o /go/src/github.com/netapp/trident/bin/tridentctl ${CLI_PKG}

	docker commit ${BUILD_CONTAINER_NAME} "sth"

	docker rm ${BUILD_CONTAINER_NAME}
	${GO_LINUX} build -ldflags $(BUILD_FLAGS) -o /go/src/github.com/netapp/trident/bin/chwrap chwrap/chwrap.go

	docker commit ${BUILD_CONTAINER_NAME} "sth"

	docker rm ${BUILD_CONTAINER_NAME}
	cp /bin/${BIN} /bin/tridentctl .
	chwrap/make-tarball.sh /bin/chwrap chwrap.tar
	docker build --build-arg PORT=${PORT} --build-arg BIN=trident_orchestrator --build-arg CLI_BIN=tridentctl --build-arg K8S=${K8S} -t ${TRIDENT_TAG} --rm .

	docker push ${TRIDENT_TAG}



tridentctl_build:
	@mkdir -p ${BIN_DIR}
	@chmod 777 ${BIN_DIR}
	@${GO_LINUX} ${BUILD} -ldflags $(BUILD_FLAGS) -o ${TRIDENT_VOLUME_PATH}/bin/${CLI_BIN} ${CLI_PKG}
ifdef CREATE_BASE_IMAGE
	@docker commit ${BUILD_CONTAINER_NAME} ${CREATE_BASE_IMAGE}
endif
	@docker rm ${BUILD_CONTAINER_NAME}
	cp ${BIN_DIR}/${CLI_BIN} .
	# docker build --build-arg PORT=${PORT} --build-arg CLI_BIN=${CLI_BIN} --build-arg K8S=${K8S} -t ${TRIDENT_TAG} --rm .
	rm ${CLI_BIN}


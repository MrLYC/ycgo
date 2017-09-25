VERSION := 0.0.1

ROOTDIR := $(shell pwd)
PROJECTNAME := ycgo

SOURCEFILES := $(shell find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./src/*" -print)

BUILDSRCDIR := ${ROOTDIR}/src
BUILDPKGDIR := ${ROOTDIR}/pkg
BUILDDIR := ${ROOTDIR}/bin
BUILDTARGET := ${BUILDDIR}/ycgo

GOPATH := ${ROOTDIR}
GO15VENDOREXPERIMENT := 1
GOENV := GOPATH=${GOPATH} GO15VENDOREXPERIMENT=${GO15VENDOREXPERIMENT}

BUILDARGS := 
LDFLAGS := 
DEBUGLDFLAGS := "-n"
RELEASELDFLAGS := "-s -w"

DEBUG := 1

.PHONY: build
build: ${BUILDDIR} ${BUILDSRCDIR} ${SOURCEFILES}
	${GOENV} go build ${BUILDARGS} -i -ldflags "${DEBUGLDFLAGS}" -o ${BUILDTARGET} ${PROJECTNAME}

.PHONY: release
release: ${BUILDDIR} ${BUILDSRCDIR}
	${GOENV} DEBUG=0 go build ${BUILDARGS} -ldflags "${RELEASELDFLAGS}" -o ${BUILDTARGET} ${PROJECTNAME}

.PHONY: upx-release
upx-release: release
	upx -9 ${BUILDTARGET}

${BUILDDIR}:
	mkdir -p ${BUILDPKGDIR}
	mkdir -p ${BUILDDIR}

${BUILDSRCDIR}:
	mkdir -p ${BUILDSRCDIR}
	ln -s ${ROOTDIR} ${BUILDSRCDIR}/${PROJECTNAME}

.PHONY: check-env
check-env:
	@${GOENV} go env
	@echo ${VAULT_ENV} | sed -e 's/ /\n/g'

.PHONY: dev-init
dev-init: ${BUILDDIR} ${BUILDSRCDIR}

.PHONY: test
test:
	${GOENV} go test

.PHONY: clean
clean:
	rm -rf ${BUILDSRCDIR} || true
	rm -rf ${BUILDTARGET} || true
	rm -rf ${BUILDPKGDIR} || true

.PHONY: format
format:
	${GOENV} find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./src/*" -exec goimports -w {} \; -exec golint {} \;


.PHONY: dep build build_darwin64 build_linux64 install

# REPO_INFO is the URL of git repository.
REPO_INFO ?= $(shell git config --get remote.origin.url)

# VERSION is the git commit hash prefixed with "git-".
VERSION ?= git-$(shell git rev-parse --short HEAD)

UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)
ifeq (${UNAME_S},Linux)
	OS=linux
endif
ifeq (${UNAME_S},Darwin)
	OS=darwin
endif
ifeq (${UNAME_M},x86_64)
	ARCH=amd64
endif
ifeq ($(UNAME_M),i686)
	ARCH=386
endif

PKGROOT = github.com/Ladicle/ghctl
OUT = ./build

dep:
	dep ensure -update

build:
	CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH} \
	go build -ldflags "-w -X $(PKGROOT)/cmd.version=$(VERSION) -X $(PKGROOT)/cmd.gitRepo=$(REPO_INFO)"

build_darwin64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
	go build -o $(OUT)/ghctl_amd64 \
           -ldflags "-w -X $(PKGROOT)/cmd.version=$(VERSION) -X $(PKGROOT)/cmd.gitRepo=$(REPO_INFO)"

build_linux64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -o $(OUT)/ghctl_linux64 \
	         -ldflags "-w -X $(PKGROOT)/cmd.version=$(VERSION) -X $(PKGROOT)/cmd.gitRepo=$(REPO_INFO)"

install:
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) \
	go install -ldflags "-w -X $(PKGROOT)/cmd.version=$(VERSION) -X $(PKGROOT)/cmd.gitRepo=$(REPO_INFO)"

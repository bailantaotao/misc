CURDIR := $(shell pwd)
GOBIN = $(shell pwd)/build/bin
DOCKER_REPOSITORY := quay.io/edwin
ifeq ($(REV),)
REV := $(shell git rev-parse --short HEAD 2> /dev/null)
endif

define my-dir
$(patsubst %/,%,$(dir $(firstword $(MAKEFILE_LIST))))
endef

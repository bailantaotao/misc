# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

PHONY += all docker test clean

include main.mk

TARGETS := $(sort $(notdir $(wildcard ./cmd/*)))
PHONY += $(TARGETS)

all: $(TARGETS)

.SECONDEXPANSION:
$(TARGETS): $(addprefix $(GOBIN)/,$$@)

$(GOBIN):
	@mkdir -p $@

$(GOBIN)/%: $(GOBIN) FORCE
	@go build -v -o $@ ./cmd/$(notdir $@)
	@echo "Done building."
	@echo "Run \"$(subst $(CURDIR),.,$@)\" to launch $(notdir $@)."

%-docker:
	@docker build -f ./cmd/$(subst -docker,,$@)/Dockerfile -t $(DOCKER_REPOSITORY)/$(subst -docker,,$@):$(REV) .
	@docker tag $(DOCKER_REPOSITORY)/$(subst -docker,,$@):$(REV) $(DOCKER_REPOSITORY)/$(subst -docker,,$@):latest

%-docker.push:
	@docker push $(DOCKER_REPOSITORY)/$(subst -docker.push,,$@):$(REV)
	@docker push $(DOCKER_REPOSITORY)/$(subst -docker.push,,$@):latest

clean:
	rm -fr $(GOBIN)/*

PHONY: help
help:
	@echo  'Generic targets:'
	@echo  '  all                         - Build all targets marked with [*]'
	@echo  '* hello                       - Build hello'
	@echo  '* log                         - Build log'
	@echo  ''
	@echo  'Docker targets:'
	@echo  '  hello-docker                 - Build hello docker image'
	@echo  '  hello-docker.push            - Push hello docker image to quay.io'
	@echo  '  log-docker                   - Build log docker image'
	@echo  '  log-docker.push              - Push log docker image to quay.io'
	@echo  '  profiling-docker             - Build profiling docker image'
	@echo  '  profiling-docker.push        - Push profiling docker image to quay.io'
	@echo  ''
	@echo  'Cleaning targets:'
	@echo  '  clean                       - Remove built executables'
	@echo  ''
	@echo  'Execute "make" or "make all" to build all targets marked with [*] '
	@echo  'For further info see the ./README.md file'

.PHONY: $(PHONY)

.PHONY: FORCE
FORCE:

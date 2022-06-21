export

PROJECT := utrade

GO_VERSION := 1.18

GO_MODULES := \
 $(CURDIR) \
 $(CURDIR)/dev

.PHONY: go-mod-tidy
go-mod-tidy: 
	@$(MAKE) -j $(GO_MODULES:%=go-mod-%-tidy)

.PHONY: $(GO_MODULES:%=go-mod-%-tidy)
$(GO_MODULES:%=go-mod-%-tidy):
	@cd $(@:go-mod-%-tidy=%) && go mod tidy -go=$(GO_VERSION)

GO_PKG_GEN := \
	internal/api

.PHONY: go-generate
go-generate:
	@$(MAKE) -j $(GO_PKG_GEN:%=go-pkg-%-gen)

.PHONY: $(GO_PKG_GEN:%=go-pkg-%-gen)
$(GO_PKG_GEN:%=go-pkg-%-gen):
	@cd $(@:go-pkg-%-gen=%) && go generate

RELEASE ?= dev

BUILD_IMAGE := $(PROJECT)/dev:$(RELEASE)

.PHONY: dev-image
dev-image:
	@echo building build image
	@docker build --build-arg go_version=$(GO_VERSION) ./dev -t $(BUILD_IMAGE)

.PHONY: dev-run
dev-run:
	@echo running the dev environment
	@docker run -it --rm \
	  -v $(CURDIR):$(CURDIR) -w $(CURDIR) \
	  -v gopkg:/go/pkg \
	  -v gocache:/home/utrade/.cache \
	  -v /var/run/docker.sock:/var/run/docker.sock \
	  $(BUILD_IMAGE) \
	  bash -c "sudo chmod o+rw /var/run/docker.sock && bash"
	  
.PHONY: functional-tests
functional-tests:
	@cd $(CURDIR)/tests && go run .
.PHONY: img build install clean test glide vendor

# go params
GOBIN = $(shell pwd)
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# binary name
BINARY_CAPTCHA=captcha

img:
	$(GOBUILD)  -o "$(GOBIN)/$(BINARY_CAPTCHA)"
	@echo "Done building."
	@echo "Run $(GOBIN)/$(BINARY_API) to launch captcha modules"

clean:
	$(GOCLEAN)
	rm -rf "$(BINARY_CAPTCHA)"
	@echo "Done cleaning."

glide:
	$(GOGET) github.com/Masterminds/glide
	$(GOINSTALL) github.com/Masterminds/glide
	@echo "Done install glide."
	@echo "Use glide --help to see help."

vendor:
	@echo "下载 golang 依赖"
	glide install

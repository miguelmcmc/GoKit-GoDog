.PHONY: clean kit run

# Build variables
AUTHOR := $(shell git log -1 --pretty=format:'%ae')
DATE := $(shell date +'%d.%m.%Y@%H:%M:%S')
GIT_COMMIT := $(shell git rev-list -1 HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
CLIENT_NAME = kydom-kit
VErsION := 0.0.1 alpha
PACKAGE := ironchip.net/kydom/kit/
CPACKAGE := $(CPACKAGE)+"vErsion"

# Detect operating system
ifeq ($(OS),Windows_NT)
	GOOS += windows
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		GOARCH += amd64
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),x86)
		GOARCH += 386
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		GOOS += linux
	endif
	ifeq ($(UNAME_S),Darwin)
		GOOS += darwin
	endif
		UNAME_P := $(shell uname -p)
	ifeq ($(UNAME_P),x86_64)
		GOARCH += amd64
	endif
		ifneq ($(filter %86,$(UNAME_P)),)
		GOARCH += 386
		endif
	ifneq ($(filter arm%,$(UNAME_P)),)
		GOARCH +=arm
	endif
endif

all: clean kit

kit:
	@echo -n "Building kit..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X $(CPACKAGE).ProgramName=$(CLIENT_NAME) -X $(CPACKAGE).GitCommit=$(GIT_COMMIT) -X $(CPACKAGE).GitBranch=$(GIT_BRANCH) -X $(CPACKAGE).ProgramVersion=$(VERSION) -X $(CPACKAGE).ProgramAuthor=$(AUTHOR) -X $(CPACKAGE).ProgramBuildDate=$(DATE)" -o $(CLIENT_NAME) .
	@echo " Done!"
	@echo Executable $(CLIENT_NAME) builded successfully!

clean:
	@echo -n "Cleaning all..."
	@rm -rf $(CURDIR)/$(CLIENT_NAME)
	@echo " Done!"

doc:
	@echo -n "Documentation on: "
	@echo -e '\e]8;;http://localhost:8080/pkg/$(PACKAGE)\ahttp://localhost:8080/pkg/$(PACKAGE)\e]8;;\a'
	@godoc -http ":8080"
	@echo " Done!"

test:
	@echo -n "Testing "$(PACKAGE)
	@go test -v --godog.random
	@echo " Done!"

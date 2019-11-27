# Makefile
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_TEST_COVER=$(GO_CMD) test -cover
GO_INSTALL=$(GO_CMD) install -v
GO_DEP_DOWNLOAD=$(GO_CMD) mod download
GO_BUILD_FLAGS=-ldflags "-w -s"
# GOOS=linux

SOURCEDIR=.

all: deps test

deps:
	@echo "==>Installing dependencies ...";
	$(GO_DEP_DOWNLOAD)

test: # run tests
	@echo "==> Running tests ...";
	@$(GO_TEST_COVER) $(SOURCEDIR)/...

.PHONY: build test install clean

BINARY  := appraise
GOMOD   := github.com/ivalx1s/skill-product-appraisal/tools/appraise
INSTALL := $(HOME)/.local/bin
SRCDIR  := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	cd $(SRCDIR) && go build -o $(BINARY) .

test:
	cd $(SRCDIR) && go test ./...

install: build
	mkdir -p $(INSTALL)
	ln -sf $(SRCDIR)$(BINARY) $(INSTALL)/$(BINARY)
	@echo "Installed: $(INSTALL)/$(BINARY) -> $(SRCDIR)$(BINARY)"

clean:
	rm -f $(SRCDIR)$(BINARY)

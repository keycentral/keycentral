BUILDDIR ?= build

.PHONY: build
build:
	go build -o $(BUILDDIR)/keycentral ./tool/keycentral

.PHONY: clean
clean:
	rm -rf $(BUILDDIR)/*

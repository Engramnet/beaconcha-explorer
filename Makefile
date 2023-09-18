# light-beaconchain-explorer
EXECUTABLE=explorer
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
VERSION := $(shell git rev-parse --short HEAD)

GOLDFLAGS += -X 'github.com/pk910/light-beaconchain-explorer/utils.BuildVersion="$(VERSION)"'
GOLDFLAGS += -X 'github.com/pk910/light-beaconchain-explorer/utils.Buildtime="$(BUILDTIME)"'
GOLDFLAGS += -X 'github.com/pk910/light-beaconchain-explorer/utils.BuildRelease="$(RELEASE)"'

.PHONY: all test clean

all: test build

test:
	go test ./...

build: linux darwin
	@echo version: $(VERSION)


linux: $(LINUX)

darwin: $(DARWIN)

$(LINUX):
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -v -o bin/$(LINUX) -ldflags="-s -w $(GOLDFLAGS)" ./cmd/explorer

$(DARWIN):
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -v -o bin/$(DARWIN) -ldflags="-s -w $(GOLDFLAGS)" ./cmd/explorer

clean:
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)
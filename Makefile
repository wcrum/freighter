# Makefile for freighter

# set shell
SHELL=/bin/bash

# set go variables
GO_FILES=./...
GO_COVERPROFILE=coverage.out

# set build variables
BIN_DIRECTORY=bin
DIST_DIRECTORY=dist

# local build of freighter for current platform
# references/configuration from .goreleaser.yaml
build:
	goreleaser build --clean --snapshot --timeout 60m --single-target

# FIPS 140-3 compliant build using native Go Cryptographic Module
# Uses the validated v1.0.0 module that's included in Go 1.24+
build-fips:
	GOFIPS140=v1.0.0 go build -tags fips140 -o $(BIN_DIRECTORY)/freighter-fips ./cmd/freighter

# local build of freighter for all platforms
# references/configuration from .goreleaser.yaml
build-all:
	goreleaser build --clean --snapshot --timeout 60m

# local release of freighter for all platforms
# references/configuration from .goreleaser.yaml
release:
	goreleaser release --clean --snapshot --timeout 60m

# install depedencies
install:
	go mod tidy
	go mod download
	CGO_ENABLED=0 go install ./cmd/...

# install FIPS 140-3 compliant version
install-fips:
	go mod tidy
	go mod download
	GOFIPS140=v1.0.0 CGO_ENABLED=0 go install -tags fips140 ./cmd/...

# format go code
fmt:
	go fmt $(GO_FILES)

# vet go code
vet:
	go vet $(GO_FILES)

# test go code
test:
	go test $(GO_FILES) -cover -race -covermode=atomic -coverprofile=$(GO_COVERPROFILE)

# test FIPS 140-3 compliant version
test-fips:
	GOFIPS140=v1.0.0 go test -tags fips140 $(GO_FILES) -cover -race -covermode=atomic -coverprofile=$(GO_COVERPROFILE)

# verify FIPS 140-3 module version
verify-fips:
	GOFIPS140=v1.0.0 go version -m | grep -i fips || echo "No FIPS module info found"

# run vulnerability scan
vuln:
	govulncheck ./...

# run vulnerability scan for FIPS 140-3 version
vuln-fips:
	GOFIPS140=v1.0.0 govulncheck -tags fips140 ./...

# cleanup artifacts
clean:
	rm -rf $(BIN_DIRECTORY) $(DIST_DIRECTORY) $(GO_COVERPROFILE)

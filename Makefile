#Parameters
GOBUILD=go build
GORUN=go run
GOCLEAN=go clean
GOTEST=go test

#Binary name
BINARY_NAME=apartments


all: build
pkgs:
	$(GOBUILD) repository/*.go
	$(GOBUILD) manager/*.go
	$(GOBUILD) handler/*.go
build: test pkgs
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run: test pkgs
	$(GORUN) *.go

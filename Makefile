# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
DEPCMD=dep
DEPENSURE=$(DEPCMD) ensure
DEPUPDATE=$(DEPENSURE) -update
BINPATH=bin
BINARY=cidrtool

all: clean dep test build
build:
	    cd cidrtool && $(GOBUILD) -o ../$(BINPATH)/$(BINARY) -v
test:
		$(GOTEST) -v
clean:
		$(GOCLEAN)
		rm -f $(BINPATH)/$(BINARY)
dep:
		$(DEPENSURE)
update:
		$(DEPUPDATE)

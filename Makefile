prep:
	if test -d pkg; then rm -rf pkg; fi

self:	prep

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps deps fmt bin

fmt:
	go fmt cmd/*.go

deps:
	@GOPATH=$(shell pwd) go get -u "github.com/cooperhewitt/go-cooperhewitt-api"

bin:	self
	@GOPATH=$(shell pwd) go build -o bin/shoebox cmd/shoebox.go
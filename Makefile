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
	@GOPATH=$(shell pwd) go get -u "github.com/adrianuswarmenhoven/goconf"

bin:	self
	@GOPATH=$(shell pwd) go build -o bin/shoebox cmd/shoebox.go

yesnofix:
	curl -s -o javascript/mapzen.whosonfirst.yesnofix.js https://raw.githubusercontent.com/whosonfirst/js-mapzen-whosonfirst-yesnofix/master/src/mapzen.whosonfirst.yesnofix.js
	curl -s -o css/mapzen.whosonfirst.yesnofix.css https://raw.githubusercontent.com/whosonfirst/js-mapzen-whosonfirst-yesnofix/master/src/mapzen.whosonfirst.yesnofix.css

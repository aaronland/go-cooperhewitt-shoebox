prep:
	if test -d pkg; then rm -rf pkg; fi

self:	prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/thisisaaronland/go-cooperhewitt-shoebox
	cp -r api src/github.com/thisisaaronland/go-cooperhewitt-shoebox/
	cp -r archive src/github.com/thisisaaronland/go-cooperhewitt-shoebox/
	cp -r client src/github.com/thisisaaronland/go-cooperhewitt-shoebox/
	cp -r endpoint src/github.com/thisisaaronland/go-cooperhewitt-shoebox
	cp -r result src/github.com/thisisaaronland/go-cooperhewitt-shoebox/
	cp -r util src/github.com/thisisaaronland/go-cooperhewitt-shoebox/
	cp *.go src/github.com/thisisaaronland/go-cooperhewitt-shoebox/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	rmdeps deps fmt bin

fmt:
	go fmt cmd/*.go

deps:
	@GOPATH=$(shell pwd) go get -u "github.com/cooperhewitt/go-cooperhewitt-api"
	@GOPATH=$(shell pwd) go get -u "github.com/adrianuswarmenhoven/goconf"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

bin:	self
	@GOPATH=$(shell pwd) go build -o bin/shoebox cmd/shoebox.go

yesnofix:
	curl -s -o javascript/mapzen.whosonfirst.yesnofix.js https://raw.githubusercontent.com/whosonfirst/js-mapzen-whosonfirst-yesnofix/master/src/mapzen.whosonfirst.yesnofix.js
	curl -s -o css/mapzen.whosonfirst.yesnofix.css https://raw.githubusercontent.com/whosonfirst/js-mapzen-whosonfirst-yesnofix/master/src/mapzen.whosonfirst.yesnofix.css

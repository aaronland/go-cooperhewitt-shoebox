prep:
	if test -d pkg; then rm -rf pkg; fi

self:	prep
	if test -d src/github.com/cooperhewitt/go-cooperhewitt-api; then rm -rf src/github.com/cooperhewitt/go-cooperhewitt-api; fi
	mkdir -p src/github.com/cooperhewitt/go-cooperhewitt-api/
	cp api.go src/github.com/cooperhewitt/go-cooperhewitt-api/

deps:
	go get "github.com/jeffail/gabs"

fmt:
	go fmt *.go
	go fmt cmd/*.go

wwms:	self
	go build -o bin/wwms cmd/wwms.go

echo:	self
	go build -o bin/echo cmd/echo.go

bin: wwms echo

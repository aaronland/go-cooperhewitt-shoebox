# go-cooperhewitt-api

Go language API library for the Cooper Hewitt API.

## How this Fits

This is one of many libraries that allow anyone to easily connect their application to the Cooper Hewitt API. This one is written in Go.

## Example

Assume this is a file called [echo.go](https://github.com/cooperhewitt/go-cooperhewitt-api/blob/master/bin/echo.go):

```
package main

import (
	"flag"
	"fmt"
	"net/url"
	"org.cooperhewitt/api"
	"os"
	"strings"
)

func main() {

	token := flag.String("token", "", "token")
	flag.Parse()

	args := flag.Args()
	call := strings.Join(args, " ")

	client := api.OAuth2Client(*token)

	method := "api.test.echo"
	params := url.Values{}
	params.Set("echo", call)

	rsp, err := client.ExecuteMethod(method, &params)

	if err != nil {
		os.Exit(1)
	}

	_, api_err := rsp.Ok()

	if api_err != nil {
		os.Exit(1)
	}

	body := rsp.Body()

	var response string
	response, _ = body.Path("echo").Data().(string)

	fmt.Println(response)
	os.Exit(0)
}
```

This would yield:

```
$> echo -token ACCESS_TOKEN wub wub wub
wub wub wub
```

The `Body` method for API responses returns
[gabs.Container](https://github.com/jeffail/gabs) thingy for wrangling the
actual JSON (in all it's unknowiness) returned by the API endpoint.

## To do

* APIError blobs are not recording error codes correctly
* Setting host and endpoint in constructor
* Support for multipart-mime uploads (just because, you can't actually do those in the API)
* Better internal logging
* Proper documentation
* Update to derive most functionality from https://github.com/straup/go-flamework-api

## See also

* https://collection.cooperhewitt.org/api/
* https://github.com/jeffail/gabs

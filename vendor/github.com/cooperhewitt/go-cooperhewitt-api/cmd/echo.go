package main

import (
	"flag"
	"fmt"
	"github.com/cooperhewitt/go-cooperhewitt-api"
	"net/url"
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
		fmt.Fprintf(os.Stderr, "Failed to call %s, because '%s'\n", method, err)
		os.Exit(1)
	}

	_, api_err := rsp.Ok()

	if api_err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute %s, because '%s'\n", method, api_err.Message)
		os.Exit(1)
	}

	body := rsp.Body()

	var response string
	response, _ = body.Path("echo").Data().(string)

	fmt.Println(response)
	os.Exit(0)
}

package main

import (
	"flag"
	"fmt"
	"github.com/cooperhewitt/go-cooperhewitt-api"
	"net/url"
	"os"
)

func main() {

	token := flag.String("token", "", "token")
	flag.Parse()

	client := api.OAuth2Client(*token)

	method := "cooperhewitt.labs.whatWouldMicahSay"
	args := url.Values{}

	rsp, err := client.ExecuteMethod(method, &args)

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

	var says string
	says, _ = body.Path("micah.says").Data().(string)

	fmt.Println(says)
	os.Exit(0)
}

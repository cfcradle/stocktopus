package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	oathUrl     = "https://slack.com/api/oauth.access"
	encodedType = "application/x-www-form-urlencoded"
)

type query struct {
	code  string `json:"code"`
	state string `json:"state"`
}

func main() {

	// Temp auth code from slack
	code := os.Args[1]

	postURL, _ := url.Parse(oathUrl)
	params := url.Values{}
	params.Add("client_id", CLIENTID)
	params.Add("client_secret", CLIENTSECRET)
	params.Add("code", code)
	postURL.RawQuery = params.Encode()

	resp, err := http.Post(postURL.String(), encodedType, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: httpPost: ", err)
		return
	}

	defer resp.Body.Close()
	fmt.Print("https://github.com/thourfor/stocktopus")
}
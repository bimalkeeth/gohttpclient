package main

import (
	"fmt"
	"github.com/bimalkeeth/gohttpclient/gohttp"
	"io/ioutil"
)

var (
	githubHttpClient = gohttp.New()
)

func main() {
	fmt.Println("Hello world")
	getUrls()
}

func getUrls() {
	response, err := githubHttpClient.Get("http://api.github.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.StatusCode)
	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

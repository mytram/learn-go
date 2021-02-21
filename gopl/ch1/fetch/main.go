package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	body, ok := fetchBody(os.Args[1])

	if ok {
		fmt.Println(body)
	}
}

func fetchBody(url string) (string, bool) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)

		return "", false
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		return "", false
	}

	return string(body), true
}

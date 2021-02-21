package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fetchBody(os.Args[1])
}

func fetchBody(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)

		return false
	}

	_, body_err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	if body_err != nil {
		fmt.Fprintf(os.Stderr, "ready body: %v\n", err)
		return false
	}

	return true
}

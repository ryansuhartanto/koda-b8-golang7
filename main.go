package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ryansuhartanto/koda-b8-golang7/req"
)

func main() {
	res := req.Loading(func() (*http.Response, error) {
		return req.Characters()
	})
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	data := string(raw)

	fmt.Println(data)
}

package main

import (
	"fmt"

	"github.com/ryansuhartanto/koda-b8-golang7/req"
)

func main() {
	data := req.Loading(func() *req.CharactersData {
		return req.Characters()
	})

	fmt.Println(*data)
}

package main

import (
	"fmt"

	"github.com/ryansuhartanto/koda-b8-golang7/req"
	"github.com/ryansuhartanto/koda-b8-golang7/utils"
)

func main() {
	utils.Cls()
	data := req.Loading(func() *req.CharactersData {
		return req.Characters()
	})

	utils.Cls()
	for _, character := range data.Results {
		fmt.Printf("%s\n", character.Name)
	}
}

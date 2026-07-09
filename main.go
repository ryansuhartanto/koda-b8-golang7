package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ryansuhartanto/koda-b8-golang7/req"
	"github.com/ryansuhartanto/koda-b8-golang7/utils"
)

func main() {
	utils.Cls()
	data := req.Loading(func() *req.CharactersData {
		return req.Characters()
	})

	scanner := bufio.NewScanner(os.Stdin)
	var query string

	for query != "0" {
		utils.Cls()

		if len(query) != 0 {
			for _, character := range data.Results {
				if !strings.Contains(strings.ToLower(character.Name), query) {
					continue
				}
				fmt.Printf("%s\n", character.Name)
			}
		}

		fmt.Println()
		fmt.Print("Masukkan nama atau 0 untuk keluar: ")
		scanner.Scan()

		query = strings.TrimSpace(scanner.Text())
	}

	err := scanner.Err()
	if err != nil {
		panic(err)
	}
}

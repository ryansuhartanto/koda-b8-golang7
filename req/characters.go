package req

import (
	"encoding/json"
	"net/http"

	"github.com/ryansuhartanto/koda-b8-golang7/model"
)

type CharactersData struct {
	Info    struct{}          `json:"info"`
	Results []model.Character `json:"results"`
}

func Characters() *CharactersData {
	resp, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		panic(err)
	}

	raw := raw(resp)
	var data CharactersData

	err = json.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}

	return &data
}

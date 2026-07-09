package req

import "net/http"

func Characters() (*http.Response, error) {
	return http.Get("https://rickandmortyapi.com/api/character")
}

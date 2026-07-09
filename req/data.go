package req

import (
	"io"
	"net/http"
)

func raw(resp *http.Response) []byte {
	var raw []byte
	var err error

	raw, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return raw
}

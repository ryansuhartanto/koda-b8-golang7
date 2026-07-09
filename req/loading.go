package req

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/ryansuhartanto/koda-b8-golang7/utils"
)

func Loading(req func() (*http.Response, error)) (resp *http.Response) {
	utils.Cls()
	fmt.Println("Loading...")

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Go(func() {
		var err error
		resp, err = req()

		if err != nil {
			panic(err.Error())
		}
	})

	return
}

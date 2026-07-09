package req

import (
	"fmt"
	"sync"

	"github.com/ryansuhartanto/koda-b8-golang7/utils"
)

func Loading[T any](req func() T) (data T) {
	utils.Cls()
	fmt.Println("Loading...")

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Go(func() {
		data = req()
	})

	return
}

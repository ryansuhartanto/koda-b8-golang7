package req

import (
	"fmt"
	"sync"
)

func Loading[T any](req func() T) (data T) {
	fmt.Println("Loading...")

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Go(func() {
		data = req()
	})

	return
}

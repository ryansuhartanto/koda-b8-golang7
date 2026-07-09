package req

import (
	"fmt"
	"sync"
	"time"
)

func Loading[T any](req func() T) (data T) {
	fmt.Print("Loading data, please wait...   ")
	done := make(chan struct{})

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Go(func() {
		data = req()
	})

	// wg.Go(func() {
	// 	time.Sleep(5 * time.Second)
	// })

	go func() {
		spinnerChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				fmt.Print("\033[2D")
				fmt.Printf("%s ", spinnerChars[i])
				i = (i + 1) % len(spinnerChars)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	return
}

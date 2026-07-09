package main

import "fmt"

func Cls() {
	fmt.Print("\033c")
}

func main() {
	Cls()
	fmt.Println("Hello, World!")
}

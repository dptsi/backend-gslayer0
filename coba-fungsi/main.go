package main

import (
	"fmt"
)

func hello() {
	fmt.Println("helooo")
}

func decorate(f func()) func() {
	return func() {
		fmt.Println("ini sudah didecorate lo")
		f()
	}
}

func test() error {
	return nil
}

func main() {
	test()
}

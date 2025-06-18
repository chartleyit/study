package main

import "fmt"

func main() {
	fmt.Println("hello world")

	for i := range 10 {
		go func() {
			fmt.Printf("\nHello %d\n", i)
		}()
	}
}

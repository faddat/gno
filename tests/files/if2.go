package main

import "fmt"

func main() {
	var i int
	if i % 1000000 {
		fmt.Println("oops")
	}
}

// Error:
// 7:5: int used as bool

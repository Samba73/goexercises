package main

import "fmt"

func main() {
	//Exercise 1.1
	func(x, y int) {
		fmt.Println(x + y)
	}(2, 3)

	//Exercise 1.2
	square := func(x int) int {
		return x * x
	}
	fmt.Println(square(4))
}

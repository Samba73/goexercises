package main

import "fmt"

func main() {
	//Exercise 1.1 basic anonymous fn direct call
	func(x, y int) {
		fmt.Println(x + y)
	}(2, 3)

	//Exercise 1.2 basic anonymous fn in variable
	square := func(x int) int {
		return x * x
	}
	fmt.Println(square(4))

	//Exercise 1.3 basic vlosure

}

package main

import "fmt"

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

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

	//Exercise 1.3 basic closure

	runCounter := makeCounter()
	fmt.Println(runCounter())
	fmt.Println(runCounter())
	fmt.Println(runCounter())
}

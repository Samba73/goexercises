package main

import "fmt"

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func makeCounter2(startValue int) func() int {
	count := startValue
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

	//Exercise 2.1 basic closure

	runCounter := makeCounter()
	fmt.Println(runCounter())
	fmt.Println(runCounter())
	fmt.Println(runCounter())

	//Exercise 2.2 basic closure
	runCounter2 := makeCounter2(10)
	fmt.Println(runCounter2())
	fmt.Println(runCounter2())
	fmt.Println(runCounter2())
}

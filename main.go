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

func incrementValue() func(int) int {
	sum := 0
	return func(val int) int {
		sum += val
		return sum
	}
}

func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return (x * factor)
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

	//Exercise 3.1 basic anonymous + closure

	adder := incrementValue()
	fmt.Println(adder(5))
	fmt.Println(adder(3))
	fmt.Println(adder(10))

	//Exercise 3.2 anonymous + closure
	times2 := makeMultiplier(2)
	fmt.Println(times2(5))
	fmt.Println(times2(10))

}

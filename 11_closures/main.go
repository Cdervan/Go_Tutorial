package main

import "fmt"

// function adder contains an anon function?
// that returns an int (but the int is )
func adder() func(int) int {
	num := 0
	return func(x int) int {
		num += x
		return num
	}
}

func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

	fmt.Println(sum(8))
	fmt.Println(sum(2))
}

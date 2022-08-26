package main

import "fmt"

func main() {
	// Long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	// i = i + 1
	// 	i++
	// }

	// Short method
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }

	// FizzBuzz
	// multiple of 3: Fizz
	// multiple of 5: Buzz
	// multiple mof 3 and 5: FizzBuzz

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}

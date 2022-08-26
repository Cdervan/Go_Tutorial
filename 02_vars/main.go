package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	var name = "chron"
	// var age = 14
	var isCool = true

	// shorthand
	age := 14
	size := 5.875

	// multi
	// race := "native"
	// email := "nativeAmerican@protonmail.com"

	race, email := "native", "nativeAmerican@protonmail.com"

	fmt.Println(name, "is", age, "years old", isCool)
	// %T returns what type a var is i.e. string
	fmt.Printf("%T\n", size)
	fmt.Println(race, email)

}

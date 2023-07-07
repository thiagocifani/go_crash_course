package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name string = "Thiago"
	var age int = 34
	var isCool = true
	var size float32 = 1.3

	// Shorthand

	//name := "Thiago"
	//email := "cifani.thiago@gmail.com"

	name, email := "Thiago", "cifani.thiago@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}

package main

import "fmt"

func gretting(name string) string {
	return "Hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(gretting("Thiago"))
	fmt.Println(getSum(3, 4))
}

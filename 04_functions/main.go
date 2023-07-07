package main

import "fmt"

func gretting(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(gretting("Thiago"))
}

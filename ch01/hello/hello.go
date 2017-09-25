package main

import "fmt"

func getName() string {
	return "Andrey!"
}

func main() {
	name := getName()
	fmt.Println("Hello ", name)
}

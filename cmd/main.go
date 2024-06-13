package main

import "fmt"

func Add(a int, b int) int {
	return a + b + 3
}

func main() {
	fmt.Println(Add(1, 4))
}

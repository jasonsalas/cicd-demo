package main

import "fmt"

// Add returns the sum of a pair of integers
func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func SayHello(name string) string{
	return "hello, " + name
}

func main() {
	fmt.Println("thanks for playing!")
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	result := add(5, 6)
	result2 := addtion(1, 6)
	fmt.Println(result)
	fmt.Println(result2)

}

//calling a function
func add(x int, y int) int {
	return x + y
}

//another way of calling func

func addtion(x, y int) int {
	return x + y
}

//function returning muiltiple value
func swap(x, y string) (string, string) {
	return y, x
}

//nacked return (use only on short function)
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

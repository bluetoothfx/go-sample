package main

import (
	"fmt"
)

//package level variable
var c, python, java bool

//initialization of variables
var ice, jock int = 1, 2

func main() {
	//fmt.Println("Hello world")
	//result := add(5, 6)
	//result2 := addtion(1, 6)
	//fmt.Println(result)
	//fmt.Println(result2)
	variableExample2()
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

//variable example
func variableExample() {
	//function level variable
	var i int
	fmt.Println(i, c, python, java)
}

//variable example with initializtion
func variableExample2() {
	var c, python, java = true, false, "no!"
	fmt.Println(ice, jock, c, python, java)
}

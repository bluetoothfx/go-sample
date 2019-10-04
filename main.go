package main

import (
	"fmt"
	"math/cmplx"
)

//package level variable
var c, python, java bool

//initialization of variables
var ice, jock int = 1, 2

//bunch of variable declaration
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//fmt.Println("Hello world")
	//result := add(5, 6)
	//result2 := addtion(1, 6)
	//fmt.Println(result)
	//fmt.Println(result2)
	variablewithrange()
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
//Variables declared without an explicit initial value are given their zero value.
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

//short variable declaration
func variableExample3() {
	var i, j int = 1, 2                   // initialization with var (can be used in package level & function level)
	k := 3                                //no need to declare var (Only can be used inside a function)
	c, python, java := true, false, "no!" //multiple value initialization without var  (Only can be used inside a function)

	fmt.Println(i, j, k, c, python, java)
}

//printing format and value type
//to know more about formatting: https://golang.org/pkg/fmt/
func variablewithrange() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

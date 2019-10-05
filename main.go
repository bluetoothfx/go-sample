package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
)

//package level variable
var c, python, java bool

//initialization of variables
var ice, jock int = 1, 2

//this is  a constant
const (
	Pi = 3.14
)

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
	//constants()
	//switchCase()
	deferStacking()
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

//Type inference
func typeInference() {
	var i int
	j := i // j is an int (if j has no type then j will be type of i)

	iceream := 42           // int (this variable will be typed based on precision of the constant)
	fork := 3.142           // float64 (this variable will be typed based on precision of the constant)
	complex := 0.867 + 0.5i // complex128 (this variable will be typed based on precision of the constant)

	fmt.Println(i, j, iceream, fork, complex)
}

func constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

//Control Flow Section
//For loop example
//initialization & increment is optional
func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

//simple for loop without init & condition is
func whileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

//if else block example
func powerFunction(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

//Switch-Case example
//In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go.
//Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
func switchCase() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

//using Defer
func deferKeyword() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

//stacking Defer
//will return LIFO order
func deferStacking() {
	fmt.Println("counting")
	for index := 0; index < 10; index++ {
		defer fmt.Println(index)
	}
	fmt.Println("done!")
}

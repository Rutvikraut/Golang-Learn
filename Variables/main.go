package main

import (
	"fmt"
	"strconv"
)

var i int = 11 //variable outside the main function declare with var keyword and datatype

var (
	name    string = "Rutvik"
	surname string = "Raut" //block of variables
	mob     int    = 7020189655
)

func main() {
	fmt.Println("Variables")
	var i int
	i = 42 // calls this value of i as it is declared in the scope gives first precedence called shadowing
	fmt.Printf("%v, %T\n", i, i)
	var j float64 = 22
	fmt.Printf("%v,%T\n", j, j)
	k := 51
	fmt.Printf("%v,%T\n", k, k)

	fmt.Println(name, " ", surname)

	z := float32(i) //type conversion
	fmt.Printf("%v, %T\n", z, z)
	var x string
	x = strconv.Itoa(i) //type conversion from string to ASCII
	fmt.Printf("%v, %T", x, x)

}

// if you have a variable which is declared but not used it cause the compile time error

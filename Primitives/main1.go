package main

import "fmt"

func main() {
	fmt.Println("------- Bit Operations ---------")
	a:=10 //1010
	b:=3  //0011
	fmt.Println("Using AND Operator :",a&b) //0010
	fmt.Println("Using OR Operator : ",a|b) //1011
	fmt.Println("Using XOR Operator : ",a^b) //1001
	fmt.Println("Using AND NOT Operator : ",a&^b) //1000
	fmt.Println("Using Bitwise Left Shift Operator : ",a<<b) //0b1010000
	fmt.Println("Using Bitwise Right Shift Operator : ",a>>b) //0001
}

// & := and operator
// | := or operator
// ^ := XOR operator
// &^ := 
// << := Binary left shift
// >> := Binary right shift
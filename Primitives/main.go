package main

import "fmt"

func main() {
	fmt.Println("Primitives")

	//boolean
	var n bool //zero value , false
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 1 //check boolean
	fmt.Printf("%v, %T\n", m, m)

	//Runes

	var r rune = 'a'                       //rune is nothing but int32
	fmt.Printf("%v, %T",r,r)

}

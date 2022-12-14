package main

import "fmt"

func main() {
	fmt.Println("----- Constant -----")
	const my_const int = 42
	fmt.Printf("%v %T",my_const,my_const)
}

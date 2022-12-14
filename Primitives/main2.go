package main

import "fmt"

func main() {
	fmt.Println("------ String Slice -------")

	s := "Hello my name is Rutvik"
	b := []byte(s)
	fmt.Printf("%v %T\n", b, b)
}

//ASCII value for each letter
/* Output :
------ Bit Wise String Slice -------
[72 101 108 108 111 32 109 121 32 110 97 109 101 32 105 115 32 82 117 116 118 105 107] []uint8 */

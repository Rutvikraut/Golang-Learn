// Is it possible to change the value of the constant ?
package main

import "fmt"

const a int16 = 41

func main() {
	fmt.Println("--------------------")
	const a int = 12
	fmt.Printf("\t%v %T\n", a, a)
	fmt.Println("--------------------")
	//What will be the output ?
}

//The value in the main scope will be the output
//we can perform operations with const having same datatypes
/* for eg. 
const a int = 42
var b int= 15
fmt.Println(a+b)
then output will be 57        Try to run the code
*/

//what will happen if the datatypes are different ? -----------> It will give the error (Try with different datatypes)

/*
Eg.

const a int = 42
var a int16 = 12

It will give the error

*/

//Another example of arithmetic operation
/*
Try without declaring the datatypes
	const a =42
	var b int =12
	fmt.Println(a+b)     ----------> 54
*/
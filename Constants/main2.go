package main

import "fmt"

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )
const (
	_ = iota + 5 //0+5=5            //assign _ when you don't care about the value
	a            //1+5=6
	b            //2+5=7
	c            //3+5=8
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//you will see that iota is changing the value.
}

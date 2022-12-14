package main

import "fmt"

func main() {

	// grades := [...]int{93, 22, 10} //array declarations

	// fmt.Printf("Grades : %v", grades) //prints the array

	var grades [3]string
	grades[0]="Rutvik"
	grades[1]="Suryakant"
	grades[2]="Raut"
	fmt.Println(grades[:2])         //slices
	fmt.Println(len(grades[:2]))             //len---> return the length of the array
	a:= []int{}
	a=append(a,1,2,5)
	fmt.Println(a)
	fmt.Println(cap(a))  //return the capacity
}

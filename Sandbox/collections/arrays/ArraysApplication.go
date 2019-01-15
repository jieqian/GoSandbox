package main

import "fmt"

func main() {
	var a4 [4]int           // An array of 4 ints, initialized to all 0.
	a4[0]=1
	fmt.Println(a4[1])
	a5 := [...]int{3, 1, 5, 10, 100} // An array initialized with a fixed size of five
	fmt.Println(a5)
}
package main

import (
	"fmt"
)

func square(x *float64) float64{
	*x = *x * *x
	return *x
}

func main() {
	x := 2.0
	fmt.Println(square(&x))

	m := map[string]string{}
	m["one"] = "first"
	fmt.Println(m)
	changeMap(&m)
	changeMap2(m)
	fmt.Println(m)

}

func changeMap(mPrt *map[string]string)  {
	m := map[string]string{}
	m["two"] = "second"
	*mPrt = m
}

func changeMap2(m map[string]string)  {
	m = map[string]string{}
	m["three"] = "third"
}

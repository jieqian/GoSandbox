package main

import (
	"sort"
)

func main() {
	list := []string{}
	list = append(list,"bb")
	list = append(list,"c")
	list = append(list,"a")
	for _,str := range list {
		println(str)
	}

	println("after sorted")
	sort.Strings(list)
	for _,str := range list {
		println(str)
	}

	i := sort.SearchStrings(list, "d")
	println(i)
	//fmt.Println(i < len(list) && list[i] == "b")
}

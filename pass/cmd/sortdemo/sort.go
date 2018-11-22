package main

import (
	"sort"
	"fmt"
)

func main(){
	a := []int {1,2,3,4,35,12,54,6}

	sort.Ints(a)

	fmt.Println(a)

	for _, v:= range a{
		fmt.Println(v)
	}
}

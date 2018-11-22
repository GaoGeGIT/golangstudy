package main

import "fmt"

func main(){
	a := [10]int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(a)
	s1 := a[5: len(a) - 1]
	fmt.Println(s1)

	s2 := make([]int, 3, 10)
	fmt.Println(len(s2), cap(a))
}

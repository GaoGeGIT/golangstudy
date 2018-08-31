package main

import "fmt"

const a = iota
const  (
	b = iota
	c = iota
	_
	d = 3.14
	e = iota
)

func main(){
	fmt.Print("a的常量值为",a)
	fmt.Print("\n")
	fmt.Print("b的常量值为",b)
	fmt.Print("\n")
	fmt.Print("c的常量值为",c)
	fmt.Print("\n")
	fmt.Print("d的常量值为",d)
	fmt.Print("\n")
	fmt.Print("e的常量值为",e)
}

package main

import "fmt"

type Date struct{

}

func(self Date) String() string{
	return "data"
}

func main(){
	fmt.Println("Hello, World!")
	fmt.Println("hellow",2,"str")

	fmt.Printf("%v/n",Date{})
}

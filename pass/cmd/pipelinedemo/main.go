package main

import "fmt"
import "../../pipeline"

func main(){
	p := pipeline.ArraySource(12,344,51,61,2)

	/*
	for {
		if num, ok := <- p; ok{
			fmt.Println(num)
		}else{
			break
		}
	}
	*/
	for v := range p{
		fmt.Println(v)
	}
}

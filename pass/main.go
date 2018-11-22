package main

import "fmt"

func main(){
	var number4 = [...]int{1,2,3,4,5,6,7,8,9,10}
	slice5 := number4[4:6:6]
	fmt.Println(slice5)
}


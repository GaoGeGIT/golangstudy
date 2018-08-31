package main

type Sender chan<- int //定义了一个入通道

type Receiver <-chan int //定义了一个出通道

func main(){
	var myChannel = make(chan int, 3)
	var sender Sender = myChannel
	var receiver Receiver = myChannel

}

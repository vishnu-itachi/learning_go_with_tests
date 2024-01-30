package main

import "fmt"


const HelloPrefix  = "hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return HelloPrefix + name
}

func main(){
	fmt.Println("hello hello!!!")
}
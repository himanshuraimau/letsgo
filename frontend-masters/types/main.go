package main

import (
	"fmt"

	"enghimanshu.tech/go/types/data"
)

func main(){
	fmt.Println("Hello World")
	max := data.Instructor{
		Id: 1,
		Name: "Max",
		Department: "Computer Science",
	}
	fmt.Println(max.Print())
}
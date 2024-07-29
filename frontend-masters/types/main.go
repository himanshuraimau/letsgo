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
	kyle:= data.NewInstructor(2, "Kyle", "Computer Science")
	goCourse:= data.Cousre{Id:2, Name: "Golang", Department: "Computer Science", 
	Instructor: kyle,Duration: "5 weeks",}
	swiftOs := data.NewWorkShop("SwiftOS",max)
	fmt.Println(max.Print())
	fmt.Println(kyle.Print())
	fmt.Println(goCourse.String())
	fmt.Printf("%s",swiftOs)
}
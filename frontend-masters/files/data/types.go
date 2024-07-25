package data

import "fmt"
/*
type distance float64
type distanceToKM float64

func (miles distance) toKM() distanceToKM {
	return distanceToKM(miles * 1.60934)
}



func test() {
	var d distance
	d = 42
	km:=d.toKM()
	fmt.Print(d,km)
}




//methods 

type locatioon string

func (oring locatioon) DistanceTo(destination locatioon) distance {
	fmt.Print("Calculating distance")
	return 0
}

func testing(){
	var origin locatioon
	var destination locatioon
	value:= origin.DistanceTo(destination)
	fmt.Println(value)
}
	*/


//Structs
type User struct{
	ID int
	FirstName string
	LastName string

}
func StructTest(){
	var user User
	user.ID=42
	user.FirstName="Matt"
	user.LastName="Aimonetti"
	fmt.Println(user)
}
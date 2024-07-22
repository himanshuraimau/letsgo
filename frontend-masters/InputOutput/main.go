package main

import "fmt"


func Tax(price float32) (float32,float32) {
	return price*0.02,price*0.03
}
func TaxwithName(price float32) (stateTax float32,cityTax float32) {
	stateTax = price*0.02
	cityTax = price*0.03
	return 
}
//pointers
func birthday(age *int){
	*age++
	*age = *age+1 
	fmt.Printf("The pointer is %v the value is %v\n",age,*age)
  //	if(*age == 12){
	//	panic("Age is 12")
	//}
}

func main() {
	defer fmt.Println("This is the end")

	citytax,_ := Tax(100)
	fmt.Println(citytax)
	stateTax,cityTax := TaxwithName(100)
	fmt.Println(stateTax,cityTax)

	age := 10
	birthday(&age)
	birthday(&age)
	fmt.Println(age)

}

//Classic for
//for i:=0;i<10;i++{
//	fmt.Println(i)
//}

// for range similar to "for in" in js
//for index,value := range countries{
//	fmt.Println(index,value)
//}

// for range similar to foreach in js
//for _,value := range countries{
//	fmt.Println(value)
//}


//var endofGame = false

// for !endofGame{
// 	fmt.Println("Game is running")
// 	endofGame = true
// }
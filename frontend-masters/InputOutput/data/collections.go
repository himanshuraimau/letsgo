package data

import "fmt"

var countries [5]string
var cities []string
var codes map[string]int    // map[key_type]value_type

func init() { 
	countries[0] = "India"
	countries[1] = "USA"
	countries[2] = "UK"
	countries[3] = "Canada"
	fmt.Println(countries)
         
}
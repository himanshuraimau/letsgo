package data

import "time"

//embedding 
type Workshop struct{
	Cousre 
	Instructor
    Date time.Time
}


func NewWorkShop(name string ,instructor Instructor) Workshop {
       w:= Workshop{}
	   w.Instructor = instructor
	   return w
}



package data


// Instructor struct
type Cousre struct{   // struct name should be Course
	Id int
	Name string
	Department string
	Legacy bool
	Duration string
	Instructor Instructor
}

func (c Cousre) String() string{
	return c.Name + " is taught by " + c.Instructor.Name + " and has a duration of " + c.Duration + "." 
}

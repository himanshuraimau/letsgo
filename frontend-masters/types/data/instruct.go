package data

type Instructor struct{
	Id int
	Name string
	Department string

}

func NewInstructor(id int, name, department string) Instructor{
	return Instructor{
		Id: id,
		Name: name,
		Department: department,
	}
}
package data

import "fmt"

func (i Instructor) Print() string{
	return fmt.Sprintf("ID: %d, Name: %s, Department: %s", i.Id, i.Name, i.Department)
}
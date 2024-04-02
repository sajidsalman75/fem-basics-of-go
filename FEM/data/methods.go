package data

import "fmt"

func (instructor Instructor) GetId() int {
	return instructor.Id
}

func (instructor Instructor) Print() string {
	return fmt.Sprintf("Id: %d, First name: %s, Last name: %s, Score: %d",
		instructor.Id, instructor.FirstName, instructor.LastName, instructor.Score)
}

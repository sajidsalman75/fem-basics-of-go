package main

import (
	"fem/go/fem/data"
	"fmt"
)

func main() {
	var instructor data.Instructor
	// instructor1 := data.Instructor{Id: 1}
	instructor.FirstName = "Salman"
	instructor.LastName = "Sajid"
	instructor.Score = 10

	fmt.Println(instructor.Print())
}

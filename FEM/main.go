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

	kyle := data.NewInstructor("Salman", "Sajid")
	// fmt.Println(kyle.Print())
	// fmt.Println(instructor.Print())

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: kyle}
	// fmt.Println(goCourse)

	swiftWS := data.NewWorkshop("Swift with iOS", instructor)

	// fmt.Printf("%v", swiftWS)
	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}
}

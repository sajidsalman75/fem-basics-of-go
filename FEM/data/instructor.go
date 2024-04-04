package data

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

// factory
func NewInstructor(name string, lastname string) Instructor {
	return Instructor{FirstName: name, LastName: lastname}
}

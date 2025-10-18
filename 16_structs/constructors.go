package main

import (
	"fmt"
	"time"
)

type students struct { //making struct of order syntax is: type struct_name struct
	id        string
	name      string
	class     string
	createdAt time.Time //nanosecond
}

// making constructor -> follow abstraction, hidding the making process, initial setup goes here
func newStudent(id string, name string, class string) *students { //this function will return the pointer to the object of struct
	myStudent := students{ //making instance of the struct
		id:    id,
		name:  name,
		class: class,
	}
	fmt.Println(&myStudent, "my student address from constructor")
	return &myStudent
}

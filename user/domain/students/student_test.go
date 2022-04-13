package students

import (
	"testing"

	"github.com/pkg/errors"
)

type TestNewScenario struct {
	description  string
	name         string
	address      string
	age          int
	rollNumber   uint
	courses      []string
	studentError error
}

func TestNewStudent(t *testing.T) {
	scenarios := []TestNewScenario{
		{
			description:  "All Details are correct",
			name:         "Alpha",
			address:      "Tumkur,Karnataka",
			age:          21,
			rollNumber:   0o1,
			courses:      []string{"a", "b", "c", "d"},
			studentError: nil,
		},
		{
			description:  "Subject count is less",
			name:         "Mosaic",
			address:      "Tumkur,Karnataka",
			age:          22,
			rollNumber:   0o3,
			courses:      []string{"a", "b", "c"},
			studentError: errors.Errorf("subject count is less than 4"),
		},
		{
			description:  "All details are correct",
			name:         "Alberto",
			address:      "Tumkur,Karnataka",
			age:          23,
			rollNumber:   0o3,
			courses:      []string{"a", "b", "c", "d"},
			studentError: nil,
		},
	}

	for _, newStudent := range scenarios {
		_, err := New(newStudent.name, uint(newStudent.age), newStudent.address, newStudent.rollNumber, newStudent.courses)
		if err != nil && newStudent.studentError == nil {
			t.Errorf("For %s got %v  expected was%v", newStudent.description, err, newStudent.studentError)
		} else if err == nil && newStudent.studentError != nil {
			t.Errorf("For %s got %v  expected was%v", newStudent.description, err, newStudent.studentError)
		}
	}
}

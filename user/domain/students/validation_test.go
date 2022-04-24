package students

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/domain/courses"
)

type duplicateScenarios struct {
	description string
	student     Student
	err         error
}

func TestDuplicates(t *testing.T) {
	scenarios := []duplicateScenarios{
		{
			description: "All Courses are duplicate for the user",
			student: Student{
				Person: Person{FullName: "shivam", Age: 24, Address: "Patna"},
				Courses: []courses.Course{
					{Name: courses.CourseTypeValues()[0]},
					{Name: courses.CourseTypeValues()[0]},
					{Name: courses.CourseTypeValues()[0]},
					{Name: courses.CourseTypeValues()[0]},
				},
				RollNumber: 24,
			},
			err: errors.Errorf("duplicate courses found"),
		},
		{
			description: "Two Courses are duplicate for the user",
			student: Student{
				Person: Person{FullName: "shivam", Age: 24, Address: "Patna"},
				Courses: []courses.Course{
					{Name: courses.CourseTypeValues()[0]},
					{Name: courses.CourseTypeValues()[0]},
					{Name: courses.CourseTypeValues()[2]},
					{Name: courses.CourseTypeValues()[1]},
				},
				RollNumber: 24,
			},
			err: errors.Errorf("duplicate courses found"),
		},
	}

	for _, scenario := range scenarios {
		err := ValidateCourses(scenario.student)
		if err != nil && scenario.err == nil {
			t.Errorf("For %s got %v expected was %v", scenario.description, err, scenario.err)
		} else if err == nil && scenario.err != nil {
			t.Errorf("For %s got %v expected was %v", scenario.description, err, scenario.err)
		}
	}
}

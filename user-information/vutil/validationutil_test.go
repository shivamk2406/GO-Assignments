package vutil

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/aggregate"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/entity"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/entity/courses/enum"
)

type duplicateScenarios struct {
	description string
	student     aggregate.Student
	err         error
}

func TestDuplicates(t *testing.T) {
	scenarios := []duplicateScenarios{
		{
			description: "Duplicate Courses",
			student: aggregate.Student{Person: entity.Person{FullName: "shivam", Age: 24, Address: "Patna"},
				Courses: []entity.Course{{Name: enum.CourseValues()[0]},
					{Name: enum.CourseValues()[0]},
					{Name: enum.CourseValues()[0]},
					{Name: enum.CourseValues()[0]}},
				RollNumber: 24},
			err: errors.Errorf("duplicate courses found"),
		},
		{
			description: "Two Duplicate Courses",
			student: aggregate.Student{Person: entity.Person{FullName: "shivam", Age: 24, Address: "Patna"},
				Courses: []entity.Course{{Name: enum.CourseValues()[0]},
					{Name: enum.CourseValues()[0]},
					{Name: enum.CourseValues()[2]},
					{Name: enum.CourseValues()[1]}},
				RollNumber: 24},
			err: errors.Errorf("duplicate courses found"),
		},
	}

	for _, scenario := range scenarios {
		fmt.Println(scenario)
		err := CheckDuplicateCourses(scenario.student)
		if err != nil && scenario.err == nil {
			t.Errorf("For %s got %v expected was %v", scenario.description, err, scenario.err)
		} else if err == nil && scenario.err != nil {
			t.Errorf("For %s got %v expected was %v", scenario.description, err, scenario.err)
		}

	}

}

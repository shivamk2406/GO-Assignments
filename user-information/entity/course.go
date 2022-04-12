package entity

import (
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/entity/courses/enum"
)

type Course struct {
	Name enum.Course
}

func (name Course) validateCourse() bool {
	return name.Name.IsACourse()
}

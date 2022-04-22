package students

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/domain/courses"
)

type Student struct {
	Person
	Courses    []courses.Course `json:"courses,omitempty"`
	RollNumber uint
}

func ValidateStudent(s Student) error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.Address, validation.Required, validation.Length(5, 100)),
		validation.Field(&s.Age, validation.By(checkNegativeValue)),
		validation.Field(&s.RollNumber, validation.By(checkNegativeValue)),
		validation.Field(&s.FullName, validation.Required, validation.Length(5, 100)),
		validation.Field(&s.Courses, validation.Required, validation.Length(4, 6)),
	)
}

func checkNegativeValue(value interface{}) error {
	s, _ := value.(int)
	if s < 0 {
		return errors.Errorf("negative value")
	}
	return nil
}

func New(name string, age uint, address string, rollNumber uint, newcourses []string) (Student, error) {
	var student Student

	student.Person = Person{FullName: name, Age: age, Address: address}
	student.RollNumber = rollNumber
	for i := 0; i < len(newcourses); i++ {
		course, err := courses.CourseTypeString(newcourses[i])
		if err != nil {
			return Student{}, err
		}

		student.Courses = append(student.Courses, courses.Course{Name: course})
	}

	err := ValidateStudent(student)
	if err != nil {
		return Student{}, err
	}

	return student, nil
}
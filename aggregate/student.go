package aggregate

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/entity"
	enum "github.com/shivamk2406/GO-Assignments/tree/Assignment-2/entity/courses/enum"
)

type Student struct {
	person  entity.Person
	courses []entity.Course
}

func (s Student) validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.person.Address, validation.Required, validation.Length(5, 100)),
		validation.Field(&s.person.Age, validation.By(checkNegativeValue)),
		validation.Field(&s.person.FullName, validation.Required, validation.Length(5, 100)),
		validation.Field(&s.courses, validation.Required, validation.Length(4, 6)),
	)
}

func checkNegativeValue(value interface{}) error {
	s, _ := value.(int)
	if s < 0 {
		return errors.Errorf("negative value")
	}
	return nil
}
func New(name string, age uint, address string, courses []string) (Student, error) {
	var student Student
	var err error
	var course entity.

	student.person = entity.Person{FullName: name, Age: age, Address: address}

	for i := 0; i < len(courses); i++ {

		course, err = enum.CourseString(courses[i])
	}

	err = student.validate()

}

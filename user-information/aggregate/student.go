package aggregate

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/entity"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/entity/courses/enum"
)

type Student struct {
	entity.Person
	Courses    []entity.Course `json:"courses,omitempty"`
	RollNumber uint
}

func validate(s Student) error {
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

func New(name string, age uint, address string, rollNumber uint, courses []string) (Student, error) {
	var student Student
	var err error

	student.Person = entity.Person{FullName: name, Age: age, Address: address}
	student.RollNumber = rollNumber
	for i := 0; i < len(courses); i++ {
		course, err := enum.CourseString(courses[i])
		if err != nil {
			return Student{}, err
		}
		student.Courses = append(student.Courses, entity.Course{Name: course})
	}
	err = validate(student)
	if err != nil {
		return Student{}, err
	}

	return student, nil
}

func (student Student) DisplayStudentDetails() {
	fmt.Printf("%s\t\t%d\t\t\t%d\t\t%s\t\t\t%v\n", student.Person.FullName, student.RollNumber, student.Age, student.Address, student.Courses)

}

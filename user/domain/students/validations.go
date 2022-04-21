package students

import (
	"github.com/pkg/errors"
)

func ValidateRollNumber(s Student) error {
	existingUsers, err := ReadFromFile()
	if err != nil {
		return err
	}

	for _, val := range existingUsers {
		if val.RollNumber == s.RollNumber {
			return errors.Errorf("user with roll number already exists")
		}
	}

	return nil
}

func ValidateCourses(s Student) error {
	visited := make(map[string]int)
	for _, course := range s.Courses {
		visited[course.Name.String()]++
		if visited[course.Name.String()] > 1 {
			return errors.Errorf("duplicate courses found")
		}
	}

	return nil
}

func Validate(s Student) error {
	err := ValidateRollNumber(s)
	if err != nil {
		return err
	}

	err = ValidateCourses(s)
	if err != nil {
		return err
	}

	return nil
}

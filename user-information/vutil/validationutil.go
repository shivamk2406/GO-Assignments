package vutil

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/aggregate"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/repository"
)

func CheckDuplicateRollNumber(s aggregate.Student) error {
	existingUsers, err := repository.ReadFromFile()
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

func CheckDuplicateCourses(s aggregate.Student) error {
	visited := make(map[string]int)
	for _, course := range s.Courses {
		visited[course.Name.String()]++
		if visited[course.Name.String()] > 1 {
			return errors.Errorf("duplicate courses found")
		}
	}
	fmt.Println(visited)
	return nil
}

func CheckDuplicates(s aggregate.Student) error {
	err := CheckDuplicateRollNumber(s)
	if err != nil {
		return err
	}
	err = CheckDuplicateCourses(s)
	if err != nil {
		return err
	}
	return nil
}

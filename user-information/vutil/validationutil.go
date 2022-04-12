package vutil

import (
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

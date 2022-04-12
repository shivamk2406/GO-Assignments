package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/pkg/errors"

	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/aggregate"
)

func ReadFromFile() ([]aggregate.Student, error) {
	data, err := ioutil.ReadFile("students.json")

	if err != nil {
		log.Println(err)
		return []aggregate.Student{}, nil
	}

	var students []aggregate.Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		log.Println(err)
		return []aggregate.Student{}, err
	}

	return students, nil
}

func AppendStudentDetails(students []aggregate.Student) error {
	existingUsers, _ := ReadFromFile()

	if students == nil {
		return errors.Errorf("no new students found")
	}
	existingUsers = append(existingUsers, students...)
	sort.Slice(existingUsers, func(i, j int) bool {
		if existingUsers[i].FullName == existingUsers[j].FullName {
			return existingUsers[i].RollNumber < existingUsers[j].RollNumber
		}
		return existingUsers[i].FullName < existingUsers[j].FullName
	})
	return SaveToFile(existingUsers)
}

func SaveToFile(students []aggregate.Student) error {
	val, err := json.MarshalIndent(students, "", " ")
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile("students.json", val, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("Updated To the File named as students.json")
	return nil
}

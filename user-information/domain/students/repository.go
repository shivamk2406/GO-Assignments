package students

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"

	"github.com/pkg/errors"
)

func ReadFromFile() ([]Student, error) {
	data, err := ioutil.ReadFile("students.json")

	if err != nil {
		log.Println(err)
		return []Student{}, nil
	}

	var newstudents []Student
	err = json.Unmarshal(data, &newstudents)
	if err != nil {
		log.Println(err)
		return []Student{}, err
	}

	return newstudents, nil
}

func AppendStudentDetails(students []Student) error {
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

func SaveToFile(students []Student) error {
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

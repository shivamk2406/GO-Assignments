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
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		log.Println(err)
		return []Student{}, nil
	}

	return decodeData(data)
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
	val, err := encodeData(students)
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(FileName, val, 0o600)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("Updated To the File named as students.json")
	return nil
}

func encodeData(students []Student) ([]byte, error) {
	val, err := json.MarshalIndent(students, "", " ")
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	return val, nil
}

func decodeData(bytes []byte) ([]Student, error) {
	var newStudents []Student
	err := json.Unmarshal(bytes, &newStudents)
	if err != nil {
		log.Println(err)
		return []Student{}, err
	}
	return newStudents, nil
}

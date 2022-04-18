package students

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func ReadFromFile() ([]Student, error) {
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		log.Println(err)
		return []Student{}, nil
	}

	return decodeData(data)
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

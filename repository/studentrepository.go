package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/aggregate"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/services"
)

func ReadFromFile() error {
	data, err := ioutil.ReadFile("students.json")

	if err != nil {
		log.Println(err)
		return err
	}

	var students []aggregate.Student
	err = json.Unmarshal(data, &students)
	if err != nil {
		log.Println(err)
		return err
	}
	services.DisplayStudentDetails(students)
	return nil
}

func SaveStudentDetails(students []aggregate.Student) error {
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
	fmt.Println("Wrote To the File named as students.json")

	return nil
}

func DeleteStudentDetails(student aggregate.Student) error {

}

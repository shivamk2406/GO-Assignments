package services

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/aggregate"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/config"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/repository"
	"golang.org/x/exp/slices"
)

func GetStudentDetails() (aggregate.Student, error) {

	var name string
	var address string
	var age uint
	var courses []string
	var extraChoice int
	var student aggregate.Student
	var rollNumber int

	fmt.Println("Enter Student Name:")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		log.Println("err")
		return aggregate.Student{}, errors.Errorf("name scanning failed")
	}

	fmt.Println("Enter Student Address:")
	_, err = fmt.Scanf("%s", &address)
	if err != nil {
		log.Println(err)
		return aggregate.Student{}, errors.Errorf("address scanning failed")
	}

	fmt.Println("Enter Student Age")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		log.Println(err)
		return aggregate.Student{}, errors.Errorf("age scanning failed")
	}

	fmt.Println("Enter Student Roll Number")
	_, err = fmt.Scanf("%d", &rollNumber)
	if err != nil {
		log.Println(err)
		return aggregate.Student{}, errors.Errorf("roll number scanning failed")
	}

	for i := 0; i < config.MaximumCourses; i++ {
		fmt.Println("Enter Course in which you want to enroll")
		var course string
		_, err = fmt.Scanf("%s", &course)
		if err != nil {
			log.Println(err)
			return aggregate.Student{}, err
		}
		courses = append(courses, course)

		if i >= config.MinimumCourses {
			fmt.Println("Minimum courses limit reached press 1 to save the existing ones as final")
			_, err := fmt.Scanf("%d", &extraChoice)
			if err != nil {
				log.Println(err)
				return aggregate.Student{}, err
			}
			if extraChoice == 1 {
				break
			}
		}

	}

	student, err = aggregate.New(name, age, address, uint(rollNumber), courses)
	if err != nil {
		return aggregate.Student{}, err
	}

	return student, nil

}

func DisplayStudentDetails(students []aggregate.Student) {
	fmt.Println("Name            Roll Number         Age              Address            Courses")
	fmt.Println("----------------------------------------------------------------------------------")
	for i := 0; i < len(students); i++ {
		students[i].DisplayStudentDetails()
	}
}

func FindStudent(rollNumber uint) int {
	students, _ := repository.ReadFromFile()
	idx := slices.IndexFunc(students, func(e aggregate.Student) bool { return e.RollNumber == rollNumber })
	return idx

}

func DeleteStudentDetails() error {
	var rollNo uint
	fmt.Println("Enter Roll Number")
	fmt.Scanf("%d", &rollNo)
	idx := FindStudent(rollNo)
	if idx == -1 {
		return errors.Errorf("no such student found")
	}
	students, _ := repository.ReadFromFile()
	students = append(students[:idx], students[idx+1:]...)
	err := repository.SaveToFile(students)
	if err != nil {
		return errors.Errorf("saving user details failed")
	}
	return nil
}

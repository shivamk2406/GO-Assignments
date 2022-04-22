package view

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/domain/students"
)

func GetStudent() (students.Student, error) {
	var name string
	var address string
	var age uint
	var courses []string
	var student students.Student
	var rollNumber int

	fmt.Println("Enter Student Name:")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		log.Println("err")
		return students.Student{}, errors.Errorf("name scanning failed")
	}

	fmt.Println("Enter Student Address:")
	_, err = fmt.Scanf("%s", &address)
	if err != nil {
		log.Println(err)
		return students.Student{}, errors.Errorf("address scanning failed")
	}

	fmt.Println("Enter Student Age")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		log.Println(err)
		return students.Student{}, errors.Errorf("age scanning failed")
	}

	fmt.Println("Enter Student Roll Number")
	_, err = fmt.Scanf("%d", &rollNumber)
	if err != nil {
		log.Println(err)
		return students.Student{}, errors.Errorf("roll number scanning failed")
	}

	courses, err = getCourses()
	if err != nil {
		return students.Student{}, err
	}

	student, err = students.New(name, age, address, uint(rollNumber), courses)
	if err != nil {
		return students.Student{}, err
	}

	err = students.ValidateStudent(student)
	if err != nil {
		return students.Student{}, err
	}

	return student, nil
}

func DisplayStudentDetails(students []students.Student) {
	fmt.Println("Name            Roll Number         Age              Address            Courses")
	fmt.Println("----------------------------------------------------------------------------------")
	for i := 0; i < len(students); i++ {
		fmt.Printf("%s\t\t%d\t\t\t%d\t\t%s\t\t\t%v\n", students[i].Person.FullName, students[i].RollNumber, students[i].Age, students[i].Address, students[i].Courses)
	}
}

func getCourses() ([]string, error) {
	var extraChoice int
	var err error
	var courses []string

	for i := 0; i < students.MaximumCourses; i++ {
		fmt.Println("Enter Course in which you want to enroll")
		var course string
		_, err = fmt.Scanf("%s", &course)
		if err != nil {
			log.Println(err)
			return []string{}, err
		}

		courses = append(courses, course)

		if i >= students.MinimumCourses {
			fmt.Println("Minimum courses limit reached press 1 to save the existing ones as final")
			_, err = fmt.Scanf("%d", &extraChoice)
			if err != nil {
				log.Println(err)
				return []string{}, err
			}
			if extraChoice == 1 {
				break
			}
		}
	}

	return courses, nil
}

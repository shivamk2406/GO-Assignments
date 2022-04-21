package students

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func GetStudent() (Student, error) {
	var name string
	var address string
	var age uint
	var courses []string
	var student Student
	var rollNumber int

	fmt.Println("Enter Student Name:")
	_, err := fmt.Scanf("%s", &name)
	if err != nil {
		log.Println("err")
		return Student{}, errors.Errorf("name scanning failed")
	}

	fmt.Println("Enter Student Address:")
	_, err = fmt.Scanf("%s", &address)
	if err != nil {
		log.Println(err)
		return Student{}, errors.Errorf("address scanning failed")
	}

	fmt.Println("Enter Student Age")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		log.Println(err)
		return Student{}, errors.Errorf("age scanning failed")
	}

	fmt.Println("Enter Student Roll Number")
	_, err = fmt.Scanf("%d", &rollNumber)
	if err != nil {
		log.Println(err)
		return Student{}, errors.Errorf("roll number scanning failed")
	}

	courses, err = getCourses()
	if err != nil {
		return Student{}, err
	}

	student, err = New(name, age, address, uint(rollNumber), courses)
	if err != nil {
		return Student{}, err
	}

	err = Validate(student)
	if err != nil {
		return Student{}, err
	}

	return student, nil
}

func DisplayStudentDetails(students []Student) {
	fmt.Println("Name            Roll Number         Age              Address            Courses")
	fmt.Println("----------------------------------------------------------------------------------")
	for i := 0; i < len(students); i++ {
		students[i].DisplayStudentDetails()
	}
}

func getCourses() ([]string, error) {
	var extraChoice int
	var err error
	var courses []string

	for i := 0; i < MaximumCourses; i++ {
		fmt.Println("Enter Course in which you want to enroll")
		var course string
		_, err = fmt.Scanf("%s", &course)
		if err != nil {
			log.Println(err)
			return []string{}, err
		}

		courses = append(courses, course)

		if i >= MinimumCourses {
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

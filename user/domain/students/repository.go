package students

import (
	"fmt"
	"log"
	"sort"

	"github.com/pkg/errors"
)

type Repository interface {
	Add(Student) error
	Load() error
	Delete(int) error
	List() ([]Student, error)
	Save() error
}

type repository struct {
	filePath string
	student  map[int]Student
}

func NewRepo() *repository {
	return &repository{}
}

func (repo *repository) Load() error {
	storedData, err := ReadFromFile()
	if err != nil {
		log.Println(err)
		return err
	}

	repo.student = make(map[int]Student)

	for _, student := range storedData {
		repo.student[int(student.RollNumber)] = student
	}
	repo.filePath = FilePath

	return nil
}

func (repo *repository) Add(student Student) error {
	err := ValidateDuplicates(student)
	if err != nil {
		log.Println(err)
		return err
	}

	if _, exists := repo.student[int(student.RollNumber)]; exists {
		return errors.Errorf("student already exists")
	}
	repo.student[int(student.RollNumber)] = student
	return nil
}

func (repo *repository) Delete(rollNumber int) error {
	if _, exists := repo.student[rollNumber]; !exists {
		return errors.Errorf("user with roll number  %d do not exists", rollNumber)
	}

	delete(repo.student, rollNumber)
	fmt.Printf("student with roll number %d deleted\n", rollNumber)
	return nil
}

func (repo *repository) Save() error {
	students := []Student{}
	for _, student := range repo.student {
		students = append(students, student)
	}

	repo.sortStudents()
	err := SaveToFile(students)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) sortStudents() []Student {
	students := []Student{}

	for _, student := range repo.student {
		students = append(students, student)
	}

	comparator := func(i, j int) bool {
		if students[i].FullName == students[j].FullName {
			return students[i].RollNumber < students[j].RollNumber
		}
		return students[i].FullName < students[j].FullName
	}

	sort.Slice(students, comparator)
	return students
}

func (repo *repository) List() ([]Student, error) {
	if len(repo.student) == 0 {
		return []Student{}, errors.Errorf("empty list nothing to display")
	}

	sortedStudents := repo.sortStudents()
	return sortedStudents, nil
}

package students

import (
	"log"
	"sort"

	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
)

type Repository interface {
	Add(Student) error
	Load() error
	Delete(int) error
	Display() error
	Save() error
}

type repository struct {
	filePath string
	student  []Student
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

	repo.student = append(repo.student, storedData...)
	repo.filePath = FileName

	return nil
}

func (repo *repository) Add(student Student) error {
	err := Validate(student)
	if err != nil {
		log.Println(err)
		return err
	}
	repo.student = append(repo.student, student)

	return nil
}

func (repo *repository) FindStudent(rollNumber uint) int {
	idx := slices.IndexFunc(repo.student, func(e Student) bool { return e.RollNumber == rollNumber })
	return idx
}

func (repo *repository) Delete(rollNumber int) error {
	idx := repo.FindStudent(uint(rollNumber))
	if idx == -1 {
		return errors.Errorf("no such student found")
	}
	repo.student = append(repo.student[:idx], repo.student[idx+1:]...)
	repo.sortStudents()
	return nil
}

func (repo *repository) Save() error {
	students := []Student{}
	for _, student := range repo.student {
		students = append(students, student)
	}

	err := SaveToFile(students)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) sortStudents() {
	sort.Slice(repo.student, func(i, j int) bool {
		if repo.student[i].FullName == repo.student[j].FullName {
			return repo.student[i].RollNumber < repo.student[j].RollNumber
		}
		return repo.student[i].FullName < repo.student[j].FullName
	})
}

func (repo *repository) Display() error {
	if len(repo.student) == 0 {
		return errors.Errorf("empty list nothing to display")
	}
	DisplayStudentDetails(repo.student)
	return nil
}

package user

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	FullName   string
	Age        uint
	Address    string
	RollNumber int
	Courses    []string
}

func GetUserName() string {
	fmt.Println("Enter Your Full Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	var name string
	scanner.Scan()
	name = scanner.Text()
	return name
}

func GetUserAge() int {
	fmt.Println("Enter Your Age:")
	scanner := bufio.NewScanner(os.Stdin)
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return int(input)
}

func GetAddress() string {
	fmt.Println("Enter Your Address:")
	scanner := bufio.NewScanner(os.Stdin)
	var address string
	scanner.Scan()
	address = scanner.Text()
	return address
}

func GetRollNumber() int {
	fmt.Println("Enter Your Roll Number")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	return int(input)
}

func GetCourses() []string {

	var courses []string
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i > 4 && i < 5; {
		fmt.Println("Enter Course in which you want to enroll!!")
		scanner.Scan()
		courses[i] = scanner.Text()
	}
	return courses

}

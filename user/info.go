package user

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/pkg/errors"
)

type User struct {
	FullName   string   `json:"full_name,omitempty"`
	Age        uint     `json:"age,omitempty"`
	Address    string   `json:"address,omitempty"`
	RollNumber uint     `json:"roll_number,omitempty"`
	Courses    []string `json:"courses,omitempty"`
}

var TotalUsers []User

func GetUserName() string {
	fmt.Println("Enter Your Full Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	var name string
	scanner.Scan()
	name = scanner.Text()
	return name
}

func GetUserAge() uint {
	fmt.Println("Enter Your Age:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil {
		errors.Wrap(err, "blank space is not a valid age")
		log.Println(err)
	}
	return uint(input)
}

func GetAddress() string {
	fmt.Println("Enter Your Address:")
	scanner := bufio.NewScanner(os.Stdin)
	var address string
	scanner.Scan()
	address = scanner.Text()
	return address
}

func GetRollNumber() uint {
	fmt.Println("Enter Your Roll Number")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil {
		errors.Wrap(err, "blank space not a valid value to parse ")
		log.Println(err)
	}
	return uint(input)
}

func GetCourses() []string {

	var courses []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Course in which you want to enroll!!")

	for courseCount := 0; courseCount < 4; courseCount++ {
		scanner.Scan()
		courses = append(courses, scanner.Text())
		fmt.Printf("%d enrolled %d left\n", courseCount+1, 3-courseCount)
		// courseCount++
	}
	return courses

}

func DisplayUserDetails() {
	fmt.Println("User Details are: ")
	fmt.Println("Name\tAge\tRoll Number\tAddress\tCourse")
	for _, value := range TotalUsers {
		fmt.Printf("%s\t%d\t%d\t%s\t", value.FullName, value.Age, value.RollNumber, value.Address)
		fmt.Println(value.Courses)

	}
	/*fmt.Println("Full Name ", user.FullName)
	fmt.Println("Age", user.Age)
	fmt.Println("Address", user.Address)
	fmt.Println("Courses", user.Courses)
	fmt.Println("Roll Number", user.RollNumber)*/
}

func SaveUserDetails(user User) {
	TotalUsers = append(TotalUsers, user)
	sort.Slice(TotalUsers, func(p int, q int) bool { return TotalUsers[p].FullName < TotalUsers[q].FullName })
}

func NewUserCreated(name string, address string, rollNumber uint, age uint, courses []string) {

}

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
	RollNumber uint
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

func GetRollNumber() uint {
	fmt.Println("Enter Your Roll Number")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.ParseUint(scanner.Text(), 10, 64)
	return uint(input)
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

func DisplayUserDetails(user User)
{
	fmt.Println("User Details are: ")
	fmt.Println("Full Name ",user.FullName)
	fmt.Println("Age",user.Age)
	fmt.Println("Address",user.Address)
	fmt.Println("Courses",user.Courses)
	fmt.Println("Roll Number",user.RollNumber)
}

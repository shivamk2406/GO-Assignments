package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func getUserInput() (pb.CreateUserRequest, error) {
	var name string
	var email string
	var subsid int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Your Name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return pb.CreateUserRequest{}, err
	}

	fmt.Println("Enter Your Email: ")
	if scanner.Scan() {
		email = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return pb.CreateUserRequest{}, err
	}

	fmt.Println("Enter Your Subsid: ")
	if scanner.Scan() {
		a, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return pb.CreateUserRequest{}, err
		}
		subsid = a
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return pb.CreateUserRequest{}, err
	}

	return pb.CreateUserRequest{Name: name, Email: email, Subsid: int32(subsid)}, nil

}

func displayChoices() {
	fmt.Println("Select Your Choice:")
	fmt.Println("1.Already a User")
	fmt.Println("2.Sign Up as a new User")
}

func getUserChoice() int {
	var choice int
	displayChoices()
	fmt.Scanf("%d", &choice)
	return choice
}

func getUserEmail() string {
	var email string
	fmt.Println("Enter Your Email: ")
	fmt.Scanf("%s", &email)
	return email
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Didn't Connect")
	}

	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	choice := getUserChoice()
	fmt.Printf("Choice is %d", choice)
	switch choice {
	case 1:
		email := getUserEmail()
		r, err := c.AuthenticateUser(ctx, &pb.AuthenticateUserRequest{Email: email})
		if err != nil {
			log.Println(err)
			log.Println("could not create user")
		}
		log.Printf("%v", r)
	case 2:
		user, err := getUserInput()
		if err != nil {
			log.Println(err)
		}

		r, err := c.CreateUser(ctx, &user)
		if err != nil {
			log.Println(err)
			log.Println("could not create user")
		}
		log.Printf("%d", r.Subsid)
	}

}

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	newspb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/news"
	subspb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/subscriptions"
	userpb "github.com/shivamk2406/newsletter-subscriptions/internal/proto/user"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func getUserInput() (userpb.CreateUserRequest, error) {
	var name string
	var email string
	//var subsid int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Your Name: ")
	if scanner.Scan() {
		name = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return userpb.CreateUserRequest{}, err
	}

	fmt.Println("Enter Your Email: ")
	if scanner.Scan() {
		email = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return userpb.CreateUserRequest{}, err
	}

	// fmt.Println("Enter Your Subsid: ")
	// if scanner.Scan() {
	// 	a, err := strconv.Atoi(string(scanner.Bytes()))
	// 	if err != nil {
	// 		log.Println(err)
	// 		return pb.CreateUserRequest{}, err
	// 	}
	// 	subsid = a
	// }
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return userpb.CreateUserRequest{}, err
	}

	return userpb.CreateUserRequest{Name: name, Email: email}, nil

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

func GetAvailablePlans(c subspb.SubscriptionManagementServiceClient, ctx context.Context) error {
	r, err := c.ListPlans(ctx, &subspb.ListPlansRequest{})
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("The Available Plans are:")
	for _, val := range r.Subs {
		fmt.Printf("%s  |", val.Name)
		for _, val1 := range val.Genres {
			fmt.Printf("%s |", val1.Name)
		}

		fmt.Println()
	}
	return nil
}
func SetSubsciption(c subspb.SubscriptionManagementServiceClient, ctx context.Context, email string) error {
	var subsid int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Your Subsid: ")
	if scanner.Scan() {
		a, err := strconv.Atoi(string(scanner.Bytes()))
		if err != nil {
			log.Println(err)
			return err
		}
		subsid = a
	}
	r, err := c.CreateSubscription(ctx, &subspb.CreateSubscriptionRequest{Email: email, Subsid: int32(subsid)})
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func Login(u userpb.UserManagementServiceClient, s subspb.SubscriptionManagementServiceClient, n newspb.NewsServiceClient, ctx context.Context) error {
	email := getUserEmail()
	r, err := u.AuthenticateUser(ctx, &userpb.AuthenticateUserRequest{Email: email})
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(r.User.Active)
	if !r.User.Active {
		err := SetSubsciption(s, ctx, email)
		if err != nil {
			return err
		}
	}
	subsid, err := getSubs(s, ctx, email)
	if err != nil {
		return err
	}
	fmt.Println(subsid)

	res, err := n.ListNews(ctx, &newspb.ListNewsRequest{Subsid: int32(subsid)})
	if err != nil {
		return err
	}
	fmt.Println(res)

	fmt.Println("By Genre")
	res1, err := n.ListNewsByGenre(ctx, &newspb.ListNewsByGenreRequest{Genre: "Daily Brief"})
	if err != nil {
		return err
	}
	fmt.Println(res1)
	return nil
}

func getSubs(c subspb.SubscriptionManagementServiceClient, ctx context.Context, email string) (int, error) {
	r, err := c.GetSubscription(ctx, &subspb.SubscriptionRequest{Email: email})
	if err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Printf("Your Subscription Plan is %v \n", r.Name)
	fmt.Println("Your Subscription plan includes following genres: ")
	for _, val := range r.Genres {
		fmt.Println(val.Name)
	}
	return int(r.Id), nil
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Didn't Connect")
	}

	defer conn.Close()

	u := userpb.NewUserManagementServiceClient(conn)
	n := newspb.NewNewsServiceClient(conn)
	s := subspb.NewSubscriptionManagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	choice := getUserChoice()
	fmt.Printf("Choice is %d", choice)
	switch choice {
	case 1:
		err := Login(u, s, n, ctx)
		GetAvailablePlans(s, ctx)
		if err != nil {
			log.Println(err)
		}
	case 2:
		user, err := getUserInput()
		if err != nil {
			log.Println(err)
		}

		r, err := u.CreateUser(ctx, &user)
		if err != nil {
			log.Println(err)
			log.Println("could not create user")
		}
		log.Printf("%s", r.Name)
	}

}

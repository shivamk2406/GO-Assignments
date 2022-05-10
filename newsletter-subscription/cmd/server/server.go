package main

import (
	"fmt"
	"os"

	serv "github.com/shivamk2406/newsletter-subscriptions/internal/server"
)

func main() {
	if err := serv.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

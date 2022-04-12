package main

import (
	"log"

	"github.com/shivamk2406/GO-Assignments/tree/Assignment-2/view"
)

func main() {
	err := view.Initialize()
	if err != nil {
		log.Println(err)
	}
}

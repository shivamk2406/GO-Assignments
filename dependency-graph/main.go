package main

import (
	"log"

	"github.com/shivamk2406/dependency-graph/view"
)

func main() {
	err := view.Initialize()
	if err != nil {
		log.Println(err)
	}
}

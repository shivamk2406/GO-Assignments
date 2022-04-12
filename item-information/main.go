package main

import (
	"fmt"

	"github.com/shivamk2406/GO-Assignments/view"
)

func main() {
	err := view.Initialize()
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"os"

	"fhrlzmn/hacktiv8-go/assignment-1/student"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input student name in arguments!")
		fmt.Println("go run main.go <studentName>")
		os.Exit(1)
	}

	inputStudent := os.Args[1]

	find := student.Search(student.Students)
	find(inputStudent)
}

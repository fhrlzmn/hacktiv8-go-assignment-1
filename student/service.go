package student

import (
	"fmt"
	"strings"
)

func isExist(student, inputStudent string) bool {
	return strings.Contains(strings.ToLower(student), strings.ToLower(inputStudent))
}

func PrintAll(students []Student) {
	for i, v := range students {
		fmt.Println(i, v.Name)
	}
}

func Search(students []Student) func(string) {
	return func(s string) {
		for i, v := range students {
			if isExist(v.Name, s) {
				fmt.Println("ID :", i)
				v.ShowDetails()
				return
			}
		}
		fmt.Println("Student not Found!")
	}
}

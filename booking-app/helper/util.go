package helper

import (
	"fmt"
	"strings"
)

func ValidateNames(firstName string, lastName string) bool {
	var hasNamesValidLen = len(firstName) > 2 && len(lastName) > 2
	return hasNamesValidLen
}

func ValidateEmailID(email string) bool {
	return strings.Contains(email, "@")
}

func CheckForLoopWithArray() {

	input := [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, value := range input {
		fmt.Printf("Index Value %v \n", value)
	}

}

func WorkWithMap() {
	employeeSalary := make(map[string]int)
	employeeSalary["beesetti"] = 10000
	employeeSalary["balakrishna"] = 10000
	employeeSalary["venkata"] = 10000

	for _, value := range employeeSalary {
		fmt.Printf("Value %v\n", value)
	}
}

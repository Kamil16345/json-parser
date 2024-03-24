package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	executor()

}
func executor() {
	fmt.Println("Welcome in JSON parser, choose operation you'd like to do:")
	for {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("1 - Read JSON data from file and parse to struct ")
		fmt.Println("2 - Read JSON data from shell and parse to struct ")
		fmt.Println("3 - Read Go struct and parse to JSON ")
		fmt.Println("4 - Quit ")

		operation, err := chooseOperation(`^[1-4]`)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		operationInt, err := strconv.Atoi(operation)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		switch operationInt {
		case 1:
			data, err := parseFromJSONFileToStruct()
			if err != nil {
				fmt.Println("Error during parsing JSON file to struct: \n", err)
			} else {
				fmt.Println("Data from JSON file parsed successfully to struct: \n", data)
			}
		case 2:
			data, err := parseFromShellJSONToStruct()
			if err != nil {
				fmt.Println("Error during parsing JSON from shell to struct: \n", err)
			} else {
				fmt.Println("Data from cmd parsed successfully: \n", data)
			}
		case 3:
			users := &users{
				[]*user{
					{
						Name:    "Kamil",
						Address: "Łódź",
						Phone:   123123,
						Tasks: []*task{
							{
								Title:       "Go home",
								Description: "Łódź, Zgierska",
							},
							{
								Title:       "finish Go Academy",
								Description: "ASAP",
							},
						},
					},
				},
			}
			data, err := parseStructToJSON("converted_struct.json", users)
			if err != nil {
				fmt.Println("Error during parsing struct to JSON: \n", err)
			} else {
				fmt.Println("Data from struct to JSON parsed successfully: \n", data)
			}
		case 4:
			fmt.Println("Bye!")
			os.Exit(0)
		}
	}
}
func chooseOperation(regexExpression string) (string, error) {
	var operation string
	fmt.Scanln(&operation)
	regex := regexp.MustCompile(regexExpression)
	if regex.MatchString(operation) {
		return operation, nil
	}
	return "", fmt.Errorf("invalid operation, try again")
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

func parseFromJSONFileToStruct() (interface{}, error) {
	var path string
	fmt.Println("Enter your JSON file name to parse: ")
	_, err := fmt.Scanln(&path)
	if err != nil {
		_ = errors.New("error during file reading")
	}
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		_ = errors.New("error during file opening")
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	var data interface{}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON data: ", err)
		return nil, err
	}
	return data, nil
}
func parseFromShellJSONToStruct() (interface{}, error) {
	fmt.Println("Type in or paste your JSON data: ")
	var data interface{}

	err := json.NewDecoder(os.Stdin).Decode(&data)
	if err != nil {
		return nil, errors.New("error during decoding JSON string")
	}
	str := fmt.Sprintf("%v", data)
	byteData := []byte(str)
	valid := json.Valid(byteData)
	if !valid {
		errors.New("data you have provided is not in valid JSON format")
	}
	return data, nil
}
func parseStructToJSON(filepath string, data interface{}) (interface{}, error) {
	fmt.Printf("Parsing struct %T to file %s\n", data, filepath)
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		_ = errors.New("error during file opening")
	}

	structData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return nil, errors.New("error during marshalling struct")
	}
	_, err = file.Write(structData)
	return structData, nil
}

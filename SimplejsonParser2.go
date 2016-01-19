package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type mystring []interface{} // user defined type for jsonstring

var parsed mystring //storing the parsed result

func objparser(jsonData string) string {
	//Parses json objects and stores them into a map else returns the data
	fmt.Println("Inside array parser")
	var result string //stores the result to be returned
	var key string
	m := make(map[string]interface{})
	if jsonData[0] == '{' {
		jsonData = jsonData[1:] // slicing the jsondata
		for len(jsonData) > 0 {
			result, jsonData = stringParser(jsonData)
			if result != "" {
				key = result
			}
			jsonData = colonParser(jsonData)
			result, jsonData = elementParser(jsonData) // parses each element encountered
			if result != "" {
				//appending values as necessary
				i,errf := strconv.ParseFloat(result,64)
				j, erri := strconv.Atoi(result)
				if result == "null" {
					m[key] = nil
				} else if result == "true" {
					m[key] = true
				} else if result == "false" {
					m[key] = false
				} else if erri == nil {
					m[key] = j
				} else if errf == nil{
					m[key] = i
				}else {
					m[key] = result
					jsonData = commaParser(jsonData)
				}
			}
			if jsonData[0] == '}' { // end of the object
				parsed = append(parsed, m) // appending the map into the parsed slice
				return jsonData[1:]
			}
		}
	}
	return jsonData
}

func arrayParser(jsonData string) string {
	//Parses json arrays and stores them into the user defined type else returns the data
	fmt.Println("Inside array parser")
	var result string
	if jsonData[0] == '[' {
		jsonData = jsonData[1:]
		for len(jsonData) > 0 {
			result, jsonData = elementParser(jsonData) // parses each element encountered
			if result != "" {
				// appending values as necessary
				i,errf := strconv.ParseFloat(result,64)
				j,erri := strconv.Atoi(result)
				if result == "null" {
					parsed = append(parsed, nil)
				} else if result == "true" {
					parsed = append(parsed, true)
				} else if result == "false" {
					parsed = append(parsed, false)
				} else if erri == nil {
					parsed = append(parsed, j)
				} else if errf == nil{
					parsed = append(parsed,i)
				}else {
					parsed = append(parsed, result) // appending data into the slice
					jsonData = commaParser(jsonData)
				}
			}
			if jsonData[0] == ']' { //end of array
				return jsonData[1:]
			}
		}
	}
	return jsonData
}

func stringParser(jsonData string) (string, string) {
	//This function parses string elements and returns the string and remaining jsondata
	fmt.Println("Inside String parser")
	var result string                      //Default value is ""
	var index int                          // stores the index of end of string
	jsonData = strings.TrimSpace(jsonData) // Trim the data for spaces
	if jsonData[0] == '"' {                //if data starts with " its a string
		fmt.Println("string found ")
		jsonData = jsonData[1:]
		for jsonData[index] != '"' { //looping over till the end of string
			index++
		}
		result = jsonData[:index]
		fmt.Println(result)
		return result, jsonData[index+1:]
	}
	return result, jsonData
}

func booleanParser(jsonData string) (string, string) {
	// Function parses boolean elements and returns the bool value and the remaining jsondata
	fmt.Println("Inside boolean parser")
	var result string
	jsonData = strings.Trim(jsonData, " ")
	if len(jsonData) > 4 {
		//Check for true or false values
		if jsonData[0:4] == "true" {
			return "true", jsonData[4:] //slicing the jsonData
		} else if jsonData[0:5] == "false" {
			return "false", jsonData[5:]
		}
	}
	return result, jsonData

}

func colonParser(jsonData string) string {
	//Parses colon in json objects and returns the colon and the remaining jsonData
	fmt.Println("Inside colon parser")
	var result string
	jsonData = strings.Trim(jsonData, " ")

	if jsonData[0] == ':' {
		fmt.Println("colon found")
		result = string(jsonData[0])
		fmt.Println(result)
		return jsonData[1:]
	}
	return jsonData
}

func commaParser(jsonData string) string {
	//Parses comma which saperate elements and returns comma and the remaining jsonData
	var result string
	fmt.Println("inside comma parser")
	jsonData = strings.Trim(jsonData, " ")
	if jsonData[0] == ',' {
		result = ","
		fmt.Println("Comma found")
		fmt.Println(result)
		return jsonData[1:]
	}
	return jsonData

}

func nullParser(jsonData string) (string, string) {
	//Parses null values form jsonstring and returns the nil value and remaining jsonData
	fmt.Println("Inside null parser")
	var result string
	jsonData = strings.Trim(jsonData, " ")
	if len(jsonData) > 4 {
		if jsonData[0:4] == "null" {
			fmt.Println("null found")
			return "null", jsonData[4:]
		}
	}
	return result, jsonData
}
func numberParser(jsonData string) (string, string) {
	//Parses numbers in jsonstring and returns the number and remaining jsonData
	jsonData = strings.Trim(jsonData, " ")
	fmt.Println("inside number parser")
	var result [][]int
	re := regexp.MustCompile("^[-+]?[0-9]*.?[0-9]+([eE][-+]?[0-9]+)?") //init regex
	result1 := re.FindAllString(jsonData, -1)                          //Library function to find occurrence of string and returns a slice of successive elements
	if len(result1) > 0 {
		result = re.FindAllStringIndex(jsonData, -1) //Gives start and end index of the string
		fmt.Println(result)
		index := result[0][1] //index is the end  index of the number
		return result1[0], jsonData[index:]
	} else {
		return "", jsonData
	}

}
func elementParser(jsonData string) (string, string) {
	//Function tries all the parsers one by one on each element and returns the result and the remaining jsonData
	var result string
	fmt.Println("INside Element Parser")
	result, jsonData = stringParser(jsonData)
	if result != "" {
		return result, jsonData
	}
	result, jsonData = numberParser(jsonData)
	if result != "" {
		return result, jsonData
	}
	result, jsonData = booleanParser(jsonData)
	if result != "" {
		return result, jsonData
	}
	result, jsonData = nullParser(jsonData)
	if result != "" {
		return result, jsonData
	}
	jsonData = commaParser(jsonData)
	jsonData = objparser(jsonData) //These functions return only one value i.e the remaining data
	jsonData = arrayParser(jsonData)

	return result, jsonData

}

func main() {
	//Main function
	buf, err := ioutil.ReadFile("test.txt") //Reading the file completely
	if err != nil {                          //error check
		fmt.Println(err)
		os.Exit(1)
	}
	data := string(buf) //bytes to string
	//Calling the functions based on the start of the  Json string
	if data[0] == '[' {
		arrayParser(data)
	} else if data[0] == '{' {
		objparser(data)
	}
	fmt.Println(parsed)

	//traversing through the result
	/*
	for _, k := range parsed {
		fmt.Println(k)
		l := k.(map[string]interface{}) //type assertion to access map type
		for k, v := range l {
			switch vv := v.(type) { //type switching
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is number", vv)
			case bool:
				fmt.Println(k, "is bool", vv)
			case nil:
				fmt.Println(k, "is nil", vv)
			case float64:
				fmt.Println(k, "is float", vv)
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	}*/
	os.Exit(0)
}

package main

import ("fmt"
		"encoding/json"
		"io/ioutil"
		"os"
		)



var k interface{} 
func main() {
	buf,err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	//var k []data
	err = json.Unmarshal(buf,&k)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(k)
	//traversing through the result
	/*
	for _, k := range parsed {
		fmt.Println(k)
		l := k.(map[string]interface{}) //type assertion to access map type
		for k, v := range l {
			switch vv := v.(type) { //type switching
			case string:
				fmt.Println(k, "is string: ", vv)
			case int:
				fmt.Println(k, "is number: ", vv)
			case bool:
				fmt.Println(k, "is bool: ", vv)
			case nil:
				fmt.Println(k, "is nil: ", vv)
			case float64:
				fmt.Println(k, "is float: ", vv)
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	}*/
}
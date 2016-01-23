package main

import "fmt"

type json struct {
	 Bool bool
	 String string
	 Number int
	 Float float64 
	 Array []interface{}
	 Object map[string]interface{}
	 Parsed bool
	 Null interface{}
	 Type string
}

func (m json) getElement() string {
	Type := m.Type
	switch Type {
		case "String":
			return m.String
	}
	return ""
}

func main() {
	var result json
	result.String = "Underwood"
	result.Type = "String"
	fmt.Println(result.getElement())
}
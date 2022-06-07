package main

import (
"fmt"

"github.com/tidwall/gjson"
)

func main() {
	json := `{"name":{"first":"li","last":"dj"},"age":18}`
	lastName := gjson.Parse(json)
	fmt.Printf("ddd: %+v\n", lastName.Get("name"))
	fmt.Println("last name:", lastName)

	//age := gjson.Get(json, "age")
	//fmt.Println("age:", age.Int())
}

package main

import (
	"encoding/json"
	"fmt"
)

type Value struct {
	Key1 string
}

type Data struct {
	List []Value
}

type Resp struct {
	Code int32
	Msg  string
	Data Data
}

func main() {
	str := `{"code":0, "msg":"success", "data": {"list": [{"key1": "value1"}]}}`
	resp := Resp{}
	err := json.Unmarshal([]byte(str), &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("resp is %v\n", resp)
}

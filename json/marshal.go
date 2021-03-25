package main

import (
	"encoding/json"
	"fmt"
)

type Agg struct {
	//AggID interface{} `json:",string"`
	AggID int64 `json:",string"`
}

type Struct2 struct {
	F1 string `json:"f_1"`           // f_1正常解析的名字
	F2 string `json:"-"`             // -表示忽略，不解析
	F3 bool   `json:"f_3,omitempty"` // omitempty表示不解析
}

type Struct3 struct {
	Id  int64  `json:"id,string"`
	Str string `json:"str,string"`
}

var m = map[string]string{
	"m1": "m_1",
	"m2": "m_2",
}

type Struct4 struct {
	//Title string `json:"title"`
	struct5
}

type struct5 struct {
	s5 string
}

type Struct6 struct {
	S6 string `json:"s_6"`
}

func (s *Struct6) MarshalJSON() ([]byte, error) {
	//str := s.S6 + "ddd"
	str := `"` + "ddd" + `"`
	return []byte(str), nil
}

type Cn struct {
	CName string `json:"c_name"`
}

type Bn struct {
	BName string `json:"b_name"`
	Cn    `json:"cn"`
}

type An struct {
	AName string `json:"a_name"`
	Bn    `json:"bn"`
}

func main() {
	//a := &Agg{
	//	AggID: 101,
	//}
	//j, _ := json.Marshal(a)
	//fmt.Println(string(j))

	// 测试tag
	//s2 := &Struct2{
	//	F1: "f1",
	//	F2: "f2",
	//	F3: false,
	//}
	//j2, _ := json.Marshal(s2)
	//fmt.Println(string(j2))
	//
	//// 测试json tag string
	//s3 := &Struct3{Id: 100028382, Str: "str3333"}
	//j3, _ := json.Marshal(s3)
	//fmt.Println(string(j3)) // 转换输出：{"id":"100028382","str":"\"str3333\""}  当为string时，若在加string，会造成重复双引号
	//
	//j4, _ := json.Marshal(m)
	//fmt.Println(string(j4))

	//s5 := Struct4{
	//	//Title: "title",
	//	struct5{s5: "ddd"},
	//}
	//j5, _ := json.Marshal(s5)
	//fmt.Println(string(j5))

	//s66 := &Struct6{S6: "fff"}
	//j6, err := json.Marshal(s66)
	//fmt.Println(string(j6), err)

	// 测试匿名字段
	an := &An{
		AName: "aName",
		Bn: Bn{
			BName: "bName",
			Cn:    Cn{CName: "cName"},
		},
	}
	jan, err := json.Marshal(an)
	fmt.Println(string(jan), err)
}

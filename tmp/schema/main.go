package main

import (
	"fmt"
	//"github.com/gorilla/schema"
	gurl "net/url"
	"path/filepath"
	"strings"
)

type Person struct {
	Name  string
	Phone string
}



func main() {
	//values := map[string][]string{
	//	"Name":  {url.QueryEscape("大迈头条")},
	//	"Phone": {"999-999-999"},
	//}
	//person := new(Person)
	//decoder := schema.NewDecoder()
	//err := decoder.Decode(person, values)
	//fmt.Println(person, err)

	token, s1, err := URL2Token("https://pic2.zhimg.com/da8e974dc.jpg?source=e593e32d")
	fmt.Println(token, s1, err)
}

func URL2Token(url string) (token string, suffix string, err error) {
	parsedUrl, err := gurl.Parse(url)
	if err != nil {
		return
	}
	path := strings.TrimLeft(parsedUrl.Path, "/")
	suffix = strings.TrimLeft(filepath.Ext(path), ".")
	if strings.Contains(path, "/") {
		lastIndex := strings.LastIndex(path, "/")
		path = path[lastIndex+1:]
	}
	if strings.Contains(path, ".") {
		index := strings.Index(path, ".")
		path = path[:index]
	}
	if strings.Contains(path, "_") {
		index := strings.Index(path, "_")
		path = path[:index]
	}
	token = path
	return
}


package main

import (
	"fmt"
	"io"
	"net/http"
)

//func Greet(writer *bytes.Buffer, name string) {
//	//fmt.Printf("Hello, %s", name)
//	fmt.Fprintf(writer, "Hello, %s", name)
//}

// io.Writer更底层的接口
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s\n", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	//Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

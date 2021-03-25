package main

import (
	"fmt"
	_ "io"
	"net/http"
)

func imageListFunc(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "iamgeList")
	if err != nil {
		fmt.Printf("imageListErr: %s", err)
	}
}

func imageFunc(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "iamge")
	if err != nil {
		fmt.Printf("imageErr: %s", err)
	}
}

func process(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Fprintln(writer, request.Form)
}

func setCookie(writer http.ResponseWriter, request *http.Request) {
	c1 := &http.Cookie{
		Name:  "c1",
		Value: "c1_value",
	}
	c2 := &http.Cookie{
		Name:  "c2",
		Value: "c2_value",
	}
	//writer.Header().Set("Set-Cookie", c1.String())
	//writer.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(writer, c1)
	http.SetCookie(writer, c2)
}

func getCookie(writer http.ResponseWriter, request *http.Request) {
	c1, _ := request.Cookie("c1")
	cs := request.Cookies() // 获取所有的cookie，返回一个list
	fmt.Fprintln(writer, c1)
	fmt.Fprintln(writer, cs[0])
	fmt.Fprintln(writer, cs[1])
}

func main() {
	/**
	在未设置handleFunc时，会选择默认的多路复用器，访问时会输出404，如下：
	404 page not found
	*/
	//err := http.ListenAndServe("", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 使用多路复用器 && 自定义路由到不同的func
	//mux := http.NewServeMux()
	//mux.HandleFunc("/image/list/", imageListFunc)
	//mux.HandleFunc("/image/", imageFunc)
	//s := &http.Server{
	//	Addr:    "0.0.0.0:80",
	//	Handler: mux,
	//}

	// 使用默认多路复用器DefaultServeMux
	//s := &http.Server{
	//	Addr: "0.0.0.0:80",
	//}
	//http.HandleFunc("/image/list/", imageListFunc)
	//http.HandleFunc("/image/", imageFunc)
	//
	//_ = s.ListenAndServe()

	s := &http.Server{
		Addr: "0.0.0.0:80",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/setcookie", setCookie)
	http.HandleFunc("/getcookie", getCookie)
	_ = s.ListenAndServe()

}

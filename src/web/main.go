package main

import (
	"net/http"
	"fmt"
)

type fooHandler struct {}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
} 

func main() {
	// HandleFunc : 경로에 해당하는 요청이 들어왔을 때 핸들러를 등록한다.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!!")
	})

	// 핸들러를 직접 등록한다.
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar!")
	})

	// 인스턴스 형식으로 핸들러를 등록한다.
	http.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", nil)
}
package main

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"log"
)

type fooHandler struct {}

type User struct {
	FirstName  string   	`json:"first_name"`
	LastName   string   	`json:"last_name"`
	Email      string   	`json:"email"`
	CreatedAt  time.Time	`json:"created_at`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 전달 받은 json 데이터를 채워줄 인스턴스를 생성한다.
	user := new(User)

	// Body에서 json 데이터를 읽는다.
	// 데이터를 user의 형식으로 파싱한다.
	err := json.NewDecoder(r.Body).Decode(user) 

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}

	user.CreatedAt = time.Now()

	// 데이터가 변경되었고 이것을 다시 json 형식으로 encoding 한다.
	// 성공하면 byte array 값이 나옴.
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Bar"
	}

	fmt.Fprintf(w, "Hello Bar! %s", name)
}

func main() {
	port := 8000

	// 핸들러는 다음과 같이 인스턴스로 생성할 수 도 있다.
	//   => http.HandleFunc는 "DefaultServeMux"를 사용한다.
	mux := http.NewServeMux()

	// HandleFunc : 경로에 해당하는 요청이 들어왔을 때 핸들러를 등록한다.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!!")
	})

	// 핸들러를 직접 등록한다.
	mux.HandleFunc("/bar", barHandler)

	// 인스턴스 형식으로 핸들러를 등록한다.
	mux.Handle("/foo", &fooHandler{})

	// param1 : 네트워크 주소
	//   => ":3000" 사용가능한 모든 IP주소의 3000번 포트에 서버를 바인딩한다는 의미이다.
	// param2 : 라우팅 핸들러
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), mux))
}
package main

import (
	"fmt"
	"net/http"
)

func main() {
	//开放8080端口，接收http请求
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	print("code: ", code, "\n")

	//返回code对应的quote
	quote := fmt.Sprintf("quote for %s", code)
	w.Write([]byte(quote))
	//返回状态码200
	w.WriteHeader(http.StatusOK)
	return
}

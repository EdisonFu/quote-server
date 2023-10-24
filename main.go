package main

import "net/http"

func main() {
	//开放8080端口，接收http请求
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	print("code: ", code, "\n")

	return
}

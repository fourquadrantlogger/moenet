package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/index", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("index", req.Cookies())

		resp.Header().Set("Location", "http://localhost:9000/relocation")
		resp.Header().Set("Set-Cookie", "self=index")

		resp.WriteHeader(301)
	})

	http.HandleFunc("/relocation", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("relocation", req.Cookies())

		resp.Header().Set("Set-Cookie", "index=relocation")

		resp.WriteHeader(200)
	})
	http.ListenAndServe(":9000", nil)
}

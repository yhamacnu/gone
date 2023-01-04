package main

import (
	"fmt"
	"net/http"
)

type dummyHandler struct {
	msg string
}

func (dh dummyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v", dh.msg)))
}

func main() {
	http.Handle("/", &dummyHandler{msg: "gone"})
	http.ListenAndServe(":8000", nil)
}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("gone"))
// })
// http.ListenAndServe(":8000", nil)
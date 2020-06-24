package main

import (
	"net/http"
	"goprojects/cms"
)

func main() {
	http.HandleFunc("/", cms.ServerIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.ListenAndServe(":3000", nil)
}
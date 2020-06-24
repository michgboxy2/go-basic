package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	// dir, _ := os.Getwd()
	// http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
	var dir string
	port := flag.String("port", "3000", "port to server HTTP on")
	path := flag.String("path", "", "path to server")
	flag.Parse()

	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir)))
}
package main

import (
	"flag"
	"net/http"
	"log"
)

func main() {
	path := flag.String("p", ".", "Document root")
	flag.Parse()

	log.Println("Document root:", *path)
	http.Handle("/", http.FileServer(http.Dir(*path)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

func main() {
	flag.IntVar(&port, "p", 3000, "The port to run serve up your files on.")
	flag.Parse()

	fmt.Printf("Serving aces on port %v\n", port)

	err := AceServe(port)

	if err != nil {
		log.Fatalln(err)
	}
}

func AceServe(port int) error {
	host := fmt.Sprintf("localhost:%v", port)

	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}

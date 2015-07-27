package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var port int

var names = []string{
	"Serena Williams",
	"Roger Federer",
	"Andy Murray",
	"Maria Sharapova",
	"Rafael Nadal",
}

func main() {
	flag.IntVar(&port, "p", 3000, "The port to run serve up your files on.")
	flag.Parse()

	fmt.Printf("Move aside %v! We're serving up aces on port %v\n", RandName(names), port)

	err := AceServe(port)

	if err != nil {
		log.Fatalln(err)
	}
}

func AceServe(port int) error {
	host := fmt.Sprintf("localhost:%v", port)

	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}

func RandName(names []string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	return names[rand.Intn(len(names))]
}

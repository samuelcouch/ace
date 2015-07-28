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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.IntVar(&port, "p", 3000, "The port to run serve up your files on.")
}

func main() {
	flag.Parse()

	name := names[rand.Intn(len(names))]

	fmt.Printf("Move aside %v! We're serving up aces on port %v\n", name, port)

	err := AceServe(port)

	if err != nil {
		log.Fatalln(err)
	}
}

func AceServe(port int) error {
	host := fmt.Sprintf("localhost:%v", port)

	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}

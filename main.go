package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	port  int
	path  string
	names = []string{
		"Serena Williams",
		"Roger Federer",
		"Andy Murray",
		"Maria Sharapova",
		"Rafael Nadal",
	}
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.IntVar(&port, "p", 3000, "The port to run serve up your files on.")
	flag.StringVar(&path, "d", ".", "The directory to serve up")
}

func main() {
	flag.Parse()

	name := names[rand.Intn(len(names))]

	fmt.Printf("Move aside %v! We're serving up aces on port %v\n", name, port)

	err := AceServe(port, path)

	if err != nil {
		log.Fatalln(err)
	}
}

func AceServe(port int, path string) error {
	host := fmt.Sprintf("localhost:%v", port)

	return http.ListenAndServe(host, http.FileServer(http.Dir(path)))
}

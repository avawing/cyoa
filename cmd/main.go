package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// create flag for file
	port := flag.Int("port", 3000, "the port to start application on")
	filename := flag.String("file", "gopher.json", "json file containing CYOA Story")
	flag.Parse()
	// POINT to the filename
	fmt.Printf("Using file in %s \n", *filename)

	f, err := os.Open(*filename)

	story, err := cyoa.JsonStory(f)
	if err != nil {
		// not best, but we need to stop and this gets the job done for now
		panic(err)
	}

	h := cyoa.NewHandler(story, cyoa.WithTemplate(nil))
	fmt.Printf("Starting server on %d \n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

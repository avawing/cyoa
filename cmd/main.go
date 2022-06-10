package main

import (
	"cyoa"
	"flag"
	"fmt"
	"os"
)

func main() {
	// create flag for file
	filename := flag.String("file", "gopher.json", "json file containing CYOA Story")
	flag.Parse()
	// POINT to the filename
	fmt.Printf("Using file in %s", *filename)

	f, err := os.Open(*filename)

	story, err := cyoa.JsonStory(f)
	if err != nil {
		// not best, but we need to stop and this gets the job done for now
		panic(err)
	}

	// why does GoLand not allowing imports??

	fmt.Printf("%+v", story)
}

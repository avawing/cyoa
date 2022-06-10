package main

import (
	"cyoa"
	"encoding/json"
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
	if err != nil {
		// not best, but we need to stop and this gets the job done for now
		panic(err)
	}
	d := json.NewDecoder(f)

	// why does GoLand not allowing imports??
	var story cyoa.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", story)
}

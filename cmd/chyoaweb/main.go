package main

import (
	"fmt"
	"flag"
	"os"
	// "github.com/asharmm/cyoa"
	"cyoa"
	"log"
	"net/http"
)

func main() {
	
	port := flag.Int("port",3000, "The port to start the server on")

	filename := flag.String("filename", "tm_forest.json", "The JSON file with CYOA story.")
	flag.Parse()


	fmt.Println("Filename is : ", *filename)

	file, err := os.Open(*filename)

	if err != nil {
		// fmt.Println("Error opening file : ", err)
		panic(err)
	}

	story, err := cyoa.JsonStory(file); 
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Println("Starting the server on port : ", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",*port), h))
}

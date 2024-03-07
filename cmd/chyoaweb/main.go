package main

import (
	"fmt"
	"flag"
	"os"
	"json"
	"github.com/asharmm/cyoa"
)

func main() {

	filename := flag.String("filename", "tm_forest.json", "The JSON file with CYOA story.")
	flag.Parse()


	fmt.Println("Filename is : ", *filename)

	file, err := os.Open(*filename)

	if err != nil {
		// fmt.Println("Error opening file : ", err)
		panic(err)
	}


	decoder := json.NewDecoder(file)
	var story cyoa.Story
}

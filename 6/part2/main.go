package main

import (
	"log"
	"strings"

	"six/internal"
)

var lights [1000][1000]int

func main() {
	content, err := internal.ExtractFile("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	//points := []Point{}

	contentSplit := strings.Split(content, "\n")

	internal.FirstTask(contentSplit, lights)

}

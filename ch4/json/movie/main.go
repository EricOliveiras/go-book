package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title: "Casablanca",
			Year:  1942,
			Color: false,
			Actors: []string{
				"Humphrey Bogart",
				"Ingrid Bergman",
			},
		},
		{
			Title: "Bullit",
			Year:  1968,
			Color: true,
			Actors: []string{
				"Humphrey Bogart",
				"Ingrid Bergman",
			},
		},
	}

	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}

	fmt.Printf("%s", data)
}

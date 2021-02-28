package main

import "fmt"

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqeline Bisset"}},
	}
	fmt.Printf("%T %#[1]v\n", movies)
}

type Movie struct {
	Title  string
	Year   int  `json: "released"`
	Color  bool `json: "color,omitempty"`
	Actors []string
}

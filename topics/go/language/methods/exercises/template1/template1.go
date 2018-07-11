// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

type player struct {
	name string
	atBats float64
	hits float64
}
// Add imports.

// Declare a struct that represents a ball player.
// Include fields called name, atBats and hits.

// Declare a method that calculates the batting average for a player.
func (p *player ) average() float64 {
	if p.atBats == 0 {
		return 0
	}

	return p.hits / p.atBats
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	ps := []player{
		{"bill", 10, 7},
		{"jim", 12, 6},
		{"ed", 6, 4},
	}

	// Display the batting average for each player in the slice.
	for i := range ps {
		fmt.Printf("%s: AVG[.%.f]\n", ps[i].name, ps[i].average()*1000)
	}

	for _, p := range ps {
		fmt.Printf("%s: AVG[.%.f]\n", p.name, p.average()*1000)
	}
}

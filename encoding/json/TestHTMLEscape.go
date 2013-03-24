package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(group); err != nil {
		fmt.Printf("failed encoding to writer: %s", err)
	}
}

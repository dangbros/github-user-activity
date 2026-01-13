package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Username is required.")
		fmt.Println("Usage: github-user-activity <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	fmt.Printf("Fetching activity for: %s...\n", username)

}

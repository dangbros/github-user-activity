package main

import (
	"fmt"
	"os"

	"github.com/dangbros/github-user-activity/pkg/client"
	"github.com/dangbros/github-user-activity/pkg/ui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Username is required.")
		fmt.Println("Usage: github-user-activity <username>")
		os.Exit(1)
	}
	username := os.Args[1]
	fmt.Printf("Fetching activity for: %s...\n", username)
	events, err := client.FetchEvents(username)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	for _, event := range events {
		fmt.Println(ui.FormatEvent(event))
	}

	fmt.Println("end thank you")
}

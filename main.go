package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dangbros/github-user-activity/pkg/client"
	"github.com/dangbros/github-user-activity/pkg/ui"
)

func main() {
	eventType := flag.String("type", "", "filter by event type (e.g. PushEvent)")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Error!")
		os.Exit(1)
	}

	username := args[0]

	fmt.Printf("Fetching activity for: %s...\n", username)
	events, err := client.FetchEvents(username)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	filteredEvents := events
	if *eventType != "" {
		var temp []client.GithubEvent
		for _, event := range events {
			if event.Type == *eventType {
				temp = append(temp, event)
			}
		}
		filteredEvents = temp
	}

	var lines []string
	for _, event := range filteredEvents {
		lines = append(lines, ui.FormatEvent(event))
	}
	output := ui.RenderBox(lines)
	fmt.Print(output)
}

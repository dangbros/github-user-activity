package app

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/dangbros/github-user-activity/pkg/client"
	"github.com/dangbros/github-user-activity/pkg/ui"
)

type Options struct {
	Username  string
	Limit     int
	EventType string
}

func Run(args []string, out io.Writer, errOut io.Writer) error {
	options, err := ParseOptions(args)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return err
	}

	fmt.Fprintf(out, "Fetching activity for: %s...\n", options.Username)

	events, err := LoadEvents(context.Background(), options.Username, options)
	if err != nil {
		fmt.Fprintln(errOut, err)
		return err
	}

	rows := BuildRows(events, options)
	fmt.Fprintln(out, RenderOutput(options.Username, rows))

	return nil
}

func ParseOptions(args []string) (Options, error) {
	var options Options
	fs := flag.NewFlagSet("github-user-activity", flag.ContinueOnError)
	fs.IntVar(&options.Limit, "limit", 0, "limit number of events (0 = all)")
	fs.StringVar(&options.EventType, "type", "", "filter by event type (e.g. PushEvent)")

	if err := fs.Parse(args); err != nil {
		return Options{}, err
	}

	if options.Limit < 0 {
		return Options{}, fmt.Errorf("limit must be zero or greater")
	}

	remaining := fs.Args()
	if len(remaining) < 1 {
		return Options{}, fmt.Errorf("username is required")
	}

	options.Username = remaining[0]
	return options, nil
}

func LoadEvents(ctx context.Context, username string, _ Options) ([]client.GithubEvent, error) {
	return client.FetchEvents(ctx, username)
}

func BuildRows(events []client.GithubEvent, options Options) []ui.EventRow {
	filteredEvents := events
	if options.EventType != "" {
		filteredEvents = nil
		for _, event := range events {
			if event.Type == options.EventType {
				filteredEvents = append(filteredEvents, event)
			}
		}
	}

	if options.Limit > 0 && options.Limit < len(filteredEvents) {
		filteredEvents = filteredEvents[:options.Limit]
	}

	rows := make([]ui.EventRow, 0, len(filteredEvents))
	for _, event := range filteredEvents {
		rows = append(rows, ui.ToRow(event))
	}

	return rows
}

func RenderOutput(username string, rows []ui.EventRow) string {
	title := ui.RenderTitle(username)
	table := ui.RenderTable(rows)
	return title + "\n\n" + table
}

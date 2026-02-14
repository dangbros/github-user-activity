package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type GithubEvent struct {
	Type      string  `json:"type"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	CreatedAt string  `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	Action string   `json:"action"`
	Commit []Commit `json:"commits"`
}

type Commit struct {
	Message string `json:"message"`
}

func fetchPage(username string, page int) ([]GithubEvent, error) {
	url := fmt.Sprintf(
		"https://api.github.com/users/%s/events?per_page=100&page=%d",
		username,
		page,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	remaining := resp.Header.Get("X-RateLimit-Remaining")
	reset := resp.Header.Get("X-RateLimit-Reset")

	if resp.StatusCode == http.StatusForbidden && remaining == "0" {
		resetUnix, _ := strconv.ParseInt(reset, 10, 64)
		resetTime := time.Unix(resetUnix, 0).Local()

		return nil, fmt.Errorf("Github API rate limit exceeded.\n"+
			"Reset at: %s\n"+
			"Tip:set GITHUB_TOKEN env variable to increase limit.",
			resetTime.Format(time.RFC1123),
		)
	}

	if resp.StatusCode == http.StatusNotFound {
		resp.Body.Close()
		return nil, fmt.Errorf("user %s not found", username)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("github api error: status %d", resp.StatusCode)
	}

	var events []GithubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("could not decode response: %v", err)
	}

	return events, nil
}

func FetchEvents(username string) ([]GithubEvent, error) {
	var allEvents []GithubEvent

	for page := 1; page <= 10; page++ {
		events, err := fetchPage(username, page)
		if err != nil {
			if len(allEvents) > 0 {
				return allEvents, fmt.Errorf("partial results returned: %w", err)
			}
			return nil, err
		}

		if len(events) == 0 {
			break
		}

		allEvents = append(allEvents, events...)
	}

	return allEvents, nil
}

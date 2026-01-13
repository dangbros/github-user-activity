package client

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func FetchEvents(username string) ([]GithubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user %s not found", username)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api error: status %d", resp.StatusCode)
	}

	var events []GithubEvent
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return nil, fmt.Errorf("could not decode response: %v", err)
	}

	return events, nil
}

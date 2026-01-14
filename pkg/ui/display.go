package ui

import (
	"fmt"

	"github.com/dangbros/github-user-activity/pkg/client"
)

func FormatEvent(e client.GithubEvent) string {
	switch e.Type {
	case "PushEvent":
		return fmt.Sprintf("pushed commits to repo %s", e.Repo.Name)
	case "IssuesEvent":
		return fmt.Sprintf("%s an issue in repo %s", e.Payload.Action, e.Repo.Name)
	case "PullRequestEvent":
		return fmt.Sprintf("PR %s in repo %s", e.Payload.Action, e.Repo.Name)
	case "WatchEvent":
		return fmt.Sprintf("starred the repo %s", e.Repo.Name)
	case "ForkEvent":
		return fmt.Sprintf("forked the repo %s", e.Repo.Name)
	case "CreateEvent":
		return fmt.Sprintf("created repository: %s", e.Repo.Name)
	default:
		return fmt.Sprintf("performed %s in %s", e.Type, e.Repo.Name)
	}
}

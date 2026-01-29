package ui

import (
	"fmt"

	"github.com/dangbros/github-user-activity/pkg/client"
)

func FormatEvent(e client.GithubEvent) string {
	timeAgo := RelativeTime(e.CreatedAt)
	switch e.Type {
	case "PushEvent":
		return fmt.Sprintf("pushed commits to repo %s • %s", e.Repo.Name, timeAgo)
	case "IssuesEvent":
		return fmt.Sprintf("%s an issue in repo %s • %s", e.Payload.Action, e.Repo.Name, timeAgo)
	case "PullRequestEvent":
		return fmt.Sprintf("PR %s in repo %s • %s", e.Payload.Action, e.Repo.Name, timeAgo)
	case "WatchEvent":
		return fmt.Sprintf("starred the repo %s • %s", e.Repo.Name, timeAgo)
	case "ForkEvent":
		return fmt.Sprintf("forked the repo %s • %s", e.Repo.Name, timeAgo)
	case "CreateEvent":
		return fmt.Sprintf("created repository: %s • %s", e.Repo.Name, timeAgo)
	default:
		return fmt.Sprintf("performed %s in %s • %s", e.Type, e.Repo.Name, timeAgo)
	}
}

package ui

import "github.com/dangbros/github-user-activity/pkg/client"

func ToRow(e client.GithubEvent) EventRow {
	return EventRow{
		Type: shortType(e.Type),
		Repo: e.Repo.Name,
		Time: RelativeTime(e.CreatedAt),
	}
}

func shortType(t string) string {
	switch t {
	case "PushEvent":
		return "Push"
	case "WatchEvent":
		return "Star"
	case "PullRequestEvent":
		return "PR"
	case "IssuesEvent":
		return "Issue"
	case "ForkEvent":
		return "Fork"
	case "CreateEvent":
		return "Create"
	default:
		return t
	}
}

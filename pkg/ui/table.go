package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func RenderTable(rows []EventRow) string {
	//dynamic max width
	maxType := len("TYPE")
	maxRepo := len("REPO")
	maxTime := len("WHEN")

	for _, r := range rows {
		if len(r.Type) > maxType {
			maxType = len(r.Type)
		}
		if len(r.Repo) > maxRepo {
			maxRepo = len(r.Repo)
		}
		if len(r.Time) > maxTime {
			maxTime = len(r.Time)
		}
	}

	//column styles using computed widths
	typeCol := lipgloss.NewStyle().
		Width(maxType + 2).
		PaddingRight(1).
		Bold(true)

	repoCol := lipgloss.NewStyle().
		Width(maxRepo + 2).
		PaddingRight(1)

	timeCol := lipgloss.NewStyle().
		Width(maxTime + 2).
		Align(lipgloss.Right)

	//seperator row
	seperator := lipgloss.JoinHorizontal(
		lipgloss.Left,
		typeCol.
			UnsetBold().
			Foreground(lipgloss.Color("8")).
			Render(strings.Repeat("─", maxType)),

		repoCol.
			UnsetBold().
			Foreground(lipgloss.Color("8")).
			Render(strings.Repeat("─", maxType)),

		timeCol.
			UnsetBold().
			Foreground(lipgloss.Color("8")).
			Render(strings.Repeat("─", maxType)),
	)

	//header rendering
	headers := lipgloss.JoinHorizontal(
		lipgloss.Left,
		typeCol.Render("TYPE"),
		repoCol.Render("REPO"),
		timeCol.Render("WHEN"),
	)

	//render body
	var body []string

	for _, r := range rows {
		line := lipgloss.JoinHorizontal(
			lipgloss.Left,
			typeCol.Render(r.Type),
			repoCol.Render(r.Repo),
			timeCol.Render(r.Time),
		)
		body = append(body, line)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		headers,
		seperator,
		lipgloss.JoinVertical(lipgloss.Left, body...),
	)

	box := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(lipgloss.Color("8"))

	return box.Render(content)
}

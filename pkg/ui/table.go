package ui

import "github.com/charmbracelet/lipgloss"

var (
	typeCol = lipgloss.NewStyle().Width(8).PaddingRight(1).Bold(true)
	repoCol = lipgloss.NewStyle().Width(30).PaddingRight(1)
	timeCol = lipgloss.NewStyle().Width(10).Align(lipgloss.Right)
)

func RenderTable(rows []EventRow) string {
	// headerStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	// cellStyle := lipgloss.NewStyle().PaddingRight(2)
	headers := lipgloss.JoinHorizontal(
		lipgloss.Left,
		typeCol.Render("TYPE"),
		repoCol.Render("REPO"),
		timeCol.Render("WHEN"),
	)

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
		lipgloss.JoinVertical(lipgloss.Left, body...),
	)

	box := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2).BorderForeground(lipgloss.Color("8"))

	return box.Render(content)
}

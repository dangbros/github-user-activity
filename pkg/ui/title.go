package ui

import "github.com/charmbracelet/lipgloss"

func RenderTitle(username string) string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("12"))

	subtleLine := lipgloss.NewStyle().
		Foreground(lipgloss.Color("8"))

	title := titleStyle.Render("Github Activity of - " + username)
	divider := subtleLine.Render("────────────────────────────────")

	return title + "\n" + divider
}

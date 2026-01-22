package ui

import "github.com/charmbracelet/lipgloss"

func RenderBox(lines []string) string {
	boxStyle := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 1).BorderForeground(lipgloss.Color("8"))
	content := lipgloss.JoinVertical(lipgloss.Left, lines...)
	return boxStyle.Render(content)
}

package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func buildColumnStyles(w columnWidth) (typeCol, repoCol, timeCol lipgloss.Style) {
	typeCol = lipgloss.NewStyle().
		Width(w.Type + 2).
		PaddingRight(1).
		Bold(true)

	repoCol = lipgloss.NewStyle().
		Width(w.Repo + 2).
		PaddingRight(1)

	timeCol = lipgloss.NewStyle().
		Width(w.Time + 2).
		Align(lipgloss.Right)

	return typeCol, repoCol, timeCol
}

func renderHeader(typeCol, repoCol, timeCol lipgloss.Style) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		typeCol.Render("TYPE"),
		repoCol.Render("REPO"),
		timeCol.Render("WHEN"),
	)
}

func renderSeperator(typeCol, repoCol, timeCol lipgloss.Style, w columnWidth) string {
	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		typeCol.
			Foreground(lipgloss.Color("8")).
			Render(strings.Repeat("─", w.Type)),

		repoCol.
			Foreground(lipgloss.Color("8")).
			Render(strings.Repeat("─", w.Repo)),

		timeCol.
			Foreground(lipgloss.Color("8")).
			Render(strings.Repeat("─", w.Time)),
	)
}

func renderBody(rows []EventRow, typeCol, repoCol, timeCol lipgloss.Style) string {
	var body []string

	for _, r := range rows {
		styledType := styleType(r.Type)

		line := lipgloss.JoinHorizontal(
			lipgloss.Left,
			typeCol.Render(styledType),
			repoCol.Render(r.Repo),
			timeCol.Render(r.Time),
		)
		body = append(body, line)
	}

	return lipgloss.JoinVertical(lipgloss.Left, body...)
}

func wrapBox(content string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		BorderForeground(lipgloss.Color("8")).
		Render(content)
}

func RenderTable(rows []EventRow) string {
	widths := computeWidths(rows)

	typeCol, repoCol, timeCol := buildColumnStyles(widths)

	header := renderHeader(typeCol, repoCol, timeCol)

	seperator := renderSeperator(typeCol, repoCol, timeCol, widths)

	body := renderBody(rows, typeCol, repoCol, timeCol)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		seperator,
		body,
	)

	return wrapBox(content)
}

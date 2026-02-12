package ui

import "github.com/charmbracelet/lipgloss"

var typeColors = map[string]string{
	"Push":   "10", //green
	"Star":   "11", //yellow
	"Create": "14", //cyan
	"PR":     "13", //magenta
	"Issue":  "9",  //red
	"Fork":   "12", //blue
}

func styleType(t string) string {
	if color, ok := typeColors[t]; ok {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color(color)).
			Render(t)
	}

	return t
}

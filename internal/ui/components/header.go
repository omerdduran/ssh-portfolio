package components

import (
	"charm.land/lipgloss/v2"
)

var headerStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#bd93f9")).
	Background(lipgloss.Color("#44475a")).
	Padding(0, 2)

func Header(title string, width int) string {
	return headerStyle.Width(width).Render(title)
}

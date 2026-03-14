package ui

import "charm.land/lipgloss/v2"

var (
	purple = lipgloss.Color("#bd93f9")
	pink   = lipgloss.Color("#ff79c6")
	cyan   = lipgloss.Color("#8be9fd")
	gray   = lipgloss.Color("#6272a4")
	white  = lipgloss.Color("#f8f8f2")

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(purple).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(pink)

	infoStyle = lipgloss.NewStyle().
			Foreground(cyan)

	quitStyle = lipgloss.NewStyle().
			Foreground(gray).
			MarginTop(2)

	containerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(purple).
			Padding(2, 4).
			Foreground(white)
)

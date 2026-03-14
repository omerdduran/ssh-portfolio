package views

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

var (
	homeTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#bd93f9")).
			MarginBottom(1)

	homeSubtitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff79c6"))

	homeInfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8be9fd"))

	homeHintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4")).
			MarginTop(1)

	homeContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#bd93f9")).
				Padding(2, 4).
				Foreground(lipgloss.Color("#f8f8f2"))
)

func HomeView(user string, width, height int) string {
	title := homeTitleStyle.Render("Ömer Duran")
	subtitle := homeSubtitleStyle.Render("Software Engineer")
	info := homeInfoStyle.Render("Welcome to my SSH portfolio!")

	greeting := ""
	if user != "" {
		greeting = homeInfoStyle.Render(fmt.Sprintf("Hello, %s!", user))
	}

	hint := homeHintStyle.Render("Press Enter to continue  •  q to quit")

	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		subtitle,
		"",
		info,
		greeting,
		hint,
	)

	box := homeContainerStyle.Render(content)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}

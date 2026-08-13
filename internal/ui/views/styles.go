package views

import "charm.land/lipgloss/v2"

// Shared list and detail styles. These used to live in blog_list.go and
// blog_detail.go; the project views borrowed them from there, so removing the
// blog pages left them homeless. They sit here now because nothing owns them
// exclusively any more.
var (
	listTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f8f8f2"))

	listDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4")).
			PaddingLeft(6)

	listSelectedTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#50fa7b")).
				Bold(true)

	detailTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#bd93f9")).
				MarginBottom(1)

	detailDateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4"))

	detailTagStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#282a36")).
			Background(lipgloss.Color("#8be9fd")).
			Padding(0, 1)

	detailSepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#44475a"))
)

package views

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/omerdduran/ssh-portfolio/internal/ui/components"
)

var (
	menuItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f8f8f2")).
			PaddingLeft(4)

	menuSelectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#50fa7b")).
				Bold(true).
				PaddingLeft(2)

	menuCountStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4"))

	menuDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4")).
			PaddingLeft(6)
)

type MenuItem struct {
	Title string
	Desc  string
	Count int
}

func MenuView(items []MenuItem, cursor int, width, height int, activeUsers int64) string {
	header := components.Header("  Ömer Duran — Portfolio", width)

	var lines []string
	lines = append(lines, "")
	for i, item := range items {
		countStr := ""
		if item.Count > 0 {
			countStr = menuCountStyle.Render(fmt.Sprintf(" (%d)", item.Count))
		}

		if i == cursor {
			lines = append(lines, menuSelectedStyle.Render("▸ "+item.Title)+countStr)
		} else {
			lines = append(lines, menuItemStyle.Render("  "+item.Title)+countStr)
		}
		lines = append(lines, menuDescStyle.Render(item.Desc))
		lines = append(lines, "")
	}

	content := lipgloss.JoinVertical(lipgloss.Left, lines...)

	footer := components.Footer([]components.KeyHint{
		{Key: "↑/k", Desc: "up"},
		{Key: "↓/j", Desc: "down"},
		{Key: "enter", Desc: "select"},
		{Key: "q", Desc: "quit"},
	}, width, activeUsers)

	// Fill available space
	bodyHeight := height - 2 // header + footer
	body := lipgloss.NewStyle().
		Width(width).
		Height(bodyHeight).
		Render(content)

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

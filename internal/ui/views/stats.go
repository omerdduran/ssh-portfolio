package views

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/omerdduran/ssh-portfolio/internal/ui/components"
)

var (
	statsLabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4")).
			PaddingLeft(4)

	statsActiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#50fa7b")).
				Bold(true)

	statsTotalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8be9fd")).
			Bold(true)
)

func StatsView(activeUsers, totalVisitors int64, scrollOffset, width, height int) string {
	header := components.Header("  Stats", width)

	var lines []string
	lines = append(lines, "")
	lines = append(lines, statsLabelStyle.Render("Visitor Statistics"))
	lines = append(lines, "")
	lines = append(lines,
		statsLabelStyle.Render("  "+statsActiveStyle.Render(fmt.Sprintf("● Active sessions: %d", activeUsers))))
	lines = append(lines, "")
	lines = append(lines,
		statsLabelStyle.Render("  "+statsTotalStyle.Render(fmt.Sprintf("  Total visitors:  %d", totalVisitors))))
	lines = append(lines, "")
	lines = append(lines, statsLabelStyle.Render(
		lipgloss.NewStyle().Foreground(lipgloss.Color("#44475a")).Render(strings.Repeat("─", width-8))))
	lines = append(lines, "")
	lines = append(lines, statsLabelStyle.Render(
		lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4")).Render(
			"Stats are tracked in-memory and reset on server restart.")))

	contentLines := strings.Split(lipgloss.JoinVertical(lipgloss.Left, lines...), "\n")

	viewHeight := height - 2
	if scrollOffset > len(contentLines)-viewHeight {
		scrollOffset = len(contentLines) - viewHeight
	}
	if scrollOffset < 0 {
		scrollOffset = 0
	}
	end := scrollOffset + viewHeight
	if end > len(contentLines) {
		end = len(contentLines)
	}
	visible := contentLines[scrollOffset:end]

	body := lipgloss.NewStyle().
		Width(width).
		Height(viewHeight).
		Render(strings.Join(visible, "\n"))

	footer := components.Footer([]components.KeyHint{
		{Key: "↑/k", Desc: "up"},
		{Key: "↓/j", Desc: "down"},
		{Key: "esc", Desc: "back"},
		{Key: "q", Desc: "quit"},
	}, width, activeUsers)

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

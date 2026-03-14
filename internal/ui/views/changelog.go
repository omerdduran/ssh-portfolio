package views

import (
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/omerdduran/ssh-portfolio/internal/content"
	"github.com/omerdduran/ssh-portfolio/internal/ui/components"
)

var (
	clTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f8f8f2")).
			Bold(true)

	clHighlightTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#f1fa8c")).
				Bold(true)

	clDateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4"))

	clSummaryStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f8f8f2")).
			PaddingLeft(4)

	clTagStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#282a36")).
			Background(lipgloss.Color("#8be9fd")).
			Padding(0, 1)

	clSepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#44475a"))

	clMarkerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#bd93f9")).
			Bold(true)
)

func formatChangelogDate(s string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return s
	}
	return t.Format("Jan 02, 2006")
}

func ChangelogView(entries []content.ChangelogEntry, scrollOffset, width, height int) string {
	header := components.Header("  Changelog", width)

	var lines []string
	lines = append(lines, "")

	for _, e := range entries {
		date := formatChangelogDate(e.Date)

		titleStr := clTitleStyle.Render(e.Title)
		if e.Highlight {
			titleStr = clHighlightTitleStyle.Render("★ " + e.Title)
		}

		lines = append(lines, clMarkerStyle.Render("  ◆ ")+titleStr)
		lines = append(lines, "    "+clDateStyle.Render(date))

		if len(e.Tags) > 0 {
			var tags []string
			for _, t := range e.Tags {
				tags = append(tags, clTagStyle.Render(t))
			}
			lines = append(lines, "    "+strings.Join(tags, " "))
		}

		if e.Summary != "" {
			lines = append(lines, "")
			lines = append(lines, clSummaryStyle.Render(e.Summary))
		}

		lines = append(lines, "")
		lines = append(lines, "    "+clSepStyle.Render(strings.Repeat("─", width-8)))
		lines = append(lines, "")
	}

	if len(entries) == 0 {
		lines = append(lines, "    No changelog entries.")
	}

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
		{Key: "g/G", Desc: "top/bottom"},
		{Key: "esc", Desc: "back"},
		{Key: "q", Desc: "quit"},
	}, width)

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

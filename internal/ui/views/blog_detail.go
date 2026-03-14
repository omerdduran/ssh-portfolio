package views

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/omerdduran/ssh-portfolio/internal/content"
	"github.com/omerdduran/ssh-portfolio/internal/ui/components"
)

var (
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

func BlogDetailView(post content.BlogPost, scrollOffset, width, height int, activeUsers int64) string {
	header := components.Header("  Blog › "+post.Title, width)

	contentWidth := width - 4
	if contentWidth < 20 {
		contentWidth = 20
	}

	// Build detail header
	date := post.Date.Format("Jan 02, 2006")
	var tags []string
	for _, t := range post.Tags {
		tags = append(tags, detailTagStyle.Render(t))
	}
	tagLine := strings.Join(tags, " ")

	detailHeader := lipgloss.JoinVertical(lipgloss.Left,
		detailTitleStyle.Render(post.Title),
		detailDateStyle.Render(date),
		tagLine,
		detailSepStyle.Render(strings.Repeat("─", contentWidth)),
		"",
	)

	// Render markdown body
	rendered := components.RenderMarkdown(post.Body, contentWidth)

	fullContent := detailHeader + rendered
	contentLines := strings.Split(fullContent, "\n")

	// Apply scroll
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
		PaddingLeft(2).
		Render(strings.Join(visible, "\n"))

	scrollInfo := ""
	if len(contentLines) > viewHeight {
		pct := 0
		if len(contentLines)-viewHeight > 0 {
			pct = scrollOffset * 100 / (len(contentLines) - viewHeight)
		}
		scrollInfo = fmt.Sprintf(" %d%%", pct)
	}

	footer := components.Footer([]components.KeyHint{
		{Key: "↑/k", Desc: "up"},
		{Key: "↓/j", Desc: "down"},
		{Key: "g/G", Desc: "top/bottom"},
		{Key: "esc", Desc: "back"},
		{Key: "q", Desc: "quit" + scrollInfo},
	}, width, activeUsers)

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

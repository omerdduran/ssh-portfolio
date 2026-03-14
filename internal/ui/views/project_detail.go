package views

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/omerdduran/ssh-portfolio/internal/content"
	"github.com/omerdduran/ssh-portfolio/internal/ui/components"
)

var (
	projLinkStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8be9fd")).
			Underline(true)
)

func ProjectDetailView(proj content.Project, scrollOffset, width, height int, activeUsers int64) string {
	header := components.Header("  Projects › "+proj.Title, width)

	contentWidth := width - 4
	if contentWidth < 20 {
		contentWidth = 20
	}

	date := proj.Date.Format("Jan 02, 2006")

	var linkParts []string
	if proj.RepoURL != "" {
		linkParts = append(linkParts, "Repo: "+projLinkStyle.Render(proj.RepoURL))
	}
	if proj.DemoURL != "" {
		linkParts = append(linkParts, "Demo: "+projLinkStyle.Render(proj.DemoURL))
	}
	links := strings.Join(linkParts, "  ")

	var tags []string
	for _, t := range proj.Tags {
		tags = append(tags, detailTagStyle.Render(t))
	}
	tagLine := strings.Join(tags, " ")

	detailHeader := lipgloss.JoinVertical(lipgloss.Left,
		detailTitleStyle.Render(proj.Title),
		detailDateStyle.Render(date),
		links,
		tagLine,
		detailSepStyle.Render(strings.Repeat("─", contentWidth)),
		"",
	)

	rendered := components.RenderMarkdown(proj.Body, contentWidth)
	fullContent := detailHeader + rendered
	contentLines := strings.Split(fullContent, "\n")

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

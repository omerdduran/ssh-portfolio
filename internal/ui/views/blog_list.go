package views

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/omerdduran/ssh-portfolio/internal/content"
	"github.com/omerdduran/ssh-portfolio/internal/ui/components"
)

var (
	blogTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f8f8f2"))

	blogDateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4"))

	blogDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4")).
			PaddingLeft(6)

	blogSelectedTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#50fa7b")).
				Bold(true)
)

func BlogListView(posts []content.BlogPost, cursor int, width, height int, activeUsers int64) string {
	header := components.Header("  Blog", width)

	var lines []string
	lines = append(lines, "")
	for i, p := range posts {
		date := p.Date.Format("Jan 02, 2006")
		dateStr := blogDateStyle.Render(date)

		if i == cursor {
			lines = append(lines, menuSelectedStyle.Render("▸ ")+blogSelectedTitleStyle.Render(p.Title)+"  "+dateStr)
		} else {
			lines = append(lines, menuItemStyle.Render("  ")+blogTitleStyle.Render(p.Title)+"  "+dateStr)
		}

		desc := p.Description
		if len(desc) > 80 {
			desc = desc[:77] + "..."
		}
		lines = append(lines, blogDescStyle.Render(desc))
		lines = append(lines, "")
	}

	if len(posts) == 0 {
		lines = append(lines, menuItemStyle.Render("  No blog posts available."))
	}

	contentStr := lipgloss.JoinVertical(lipgloss.Left, lines...)

	footer := components.Footer([]components.KeyHint{
		{Key: "↑/k", Desc: "up"},
		{Key: "↓/j", Desc: "down"},
		{Key: "enter", Desc: "read"},
		{Key: "esc", Desc: "back"},
		{Key: "q", Desc: "quit"},
	}, width, activeUsers)

	bodyHeight := height - 2
	body := lipgloss.NewStyle().Width(width).Height(bodyHeight).Render(contentStr)

	return fmt.Sprintf("%s\n%s\n%s", header, body, footer)
}

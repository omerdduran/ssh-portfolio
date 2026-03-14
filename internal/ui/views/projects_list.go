package views

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"github.com/omerdduran/ssh-portfolio/internal/content"
	"github.com/omerdduran/ssh-portfolio/internal/ui/components"
)

var (
	projIndicatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#8be9fd"))
)

func ProjectsListView(projects []content.Project, cursor int, width, height int) string {
	header := components.Header("  Projects", width)

	var lines []string
	lines = append(lines, "")
	for i, p := range projects {
		indicators := ""
		if p.RepoURL != "" {
			indicators += projIndicatorStyle.Render(" [repo]")
		}
		if p.DemoURL != "" {
			indicators += projIndicatorStyle.Render(" [demo]")
		}

		if i == cursor {
			lines = append(lines, menuSelectedStyle.Render("▸ ")+blogSelectedTitleStyle.Render(p.Title)+indicators)
		} else {
			lines = append(lines, menuItemStyle.Render("  ")+blogTitleStyle.Render(p.Title)+indicators)
		}

		desc := p.Description
		if len(desc) > 80 {
			desc = desc[:77] + "..."
		}
		lines = append(lines, blogDescStyle.Render(desc))
		lines = append(lines, "")
	}

	if len(projects) == 0 {
		lines = append(lines, menuItemStyle.Render("  No projects available."))
	}

	contentStr := lipgloss.JoinVertical(lipgloss.Left, lines...)

	footer := components.Footer([]components.KeyHint{
		{Key: "↑/k", Desc: "up"},
		{Key: "↓/j", Desc: "down"},
		{Key: "enter", Desc: "details"},
		{Key: "esc", Desc: "back"},
		{Key: "q", Desc: "quit"},
	}, width)

	bodyHeight := height - 2
	body := lipgloss.NewStyle().Width(width).Height(bodyHeight).Render(contentStr)

	return fmt.Sprintf("%s\n%s\n%s", header, body, footer)
}

package views

import (
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/omerdduran/ssh-portfolio/internal/content"
	"github.com/omerdduran/ssh-portfolio/internal/ui/components"
)

var (
	workCompanyStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#50fa7b"))

	workRoleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff79c6"))

	workDateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4"))

	workMarkerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#bd93f9")).
			Bold(true)

	workLinkStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8be9fd")).
			Underline(true)

	workSepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#44475a"))

	workBodyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f8f8f2")).
			PaddingLeft(4)
)

func formatWorkDate(s string) string {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return s // "Current" or other string
	}
	return t.Format("Jan 2006")
}

func WorkView(entries []content.WorkEntry, scrollOffset, width, height int) string {
	header := components.Header("  Work Experience", width)

	var lines []string
	lines = append(lines, "")

	for _, w := range entries {
		start := formatWorkDate(w.DateStart)
		end := formatWorkDate(w.DateEnd)
		dateRange := start + " — " + end

		lines = append(lines, workMarkerStyle.Render("  ● ")+workCompanyStyle.Render(w.Company))
		lines = append(lines, "    "+workRoleStyle.Render(w.Role))
		lines = append(lines, "    "+workDateStyle.Render(dateRange))

		if w.CompanyURL != "" {
			lines = append(lines, "    "+workLinkStyle.Render(w.CompanyURL))
		}

		if w.Body != "" {
			lines = append(lines, "")
			// Wrap body text
			bodyLines := strings.Split(w.Body, "\n")
			for _, bl := range bodyLines {
				lines = append(lines, workBodyStyle.Render(bl))
			}
		}

		lines = append(lines, "")
		lines = append(lines, "    "+workSepStyle.Render(strings.Repeat("─", width-8)))
		lines = append(lines, "")
	}

	if len(entries) == 0 {
		lines = append(lines, "    No work experience entries.")
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

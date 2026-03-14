package components

import (
	"strings"

	"charm.land/lipgloss/v2"
)

var (
	footerBg        = lipgloss.Color("#44475a")
	footerKeyStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#f8f8f2")).Bold(true)
	footerDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4"))
	footerBarStyle  = lipgloss.NewStyle().Background(footerBg).Padding(0, 1)
)

type KeyHint struct {
	Key  string
	Desc string
}

func Footer(hints []KeyHint, width int) string {
	var parts []string
	for _, h := range hints {
		parts = append(parts, footerKeyStyle.Render(h.Key)+" "+footerDescStyle.Render(h.Desc))
	}
	bar := strings.Join(parts, "  ")
	return footerBarStyle.Width(width).Render(bar)
}

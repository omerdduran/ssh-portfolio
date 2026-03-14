package components

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

var (
	footerBg        = lipgloss.Color("#44475a")
	footerKeyStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#f8f8f2")).Bold(true)
	footerDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6272a4"))
	footerBarStyle  = lipgloss.NewStyle().Background(footerBg).Padding(0, 1)
	onlineStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#50fa7b")).Bold(true)
)

type KeyHint struct {
	Key  string
	Desc string
}

func Footer(hints []KeyHint, width int, activeUsers int64) string {
	var parts []string
	for _, h := range hints {
		parts = append(parts, footerKeyStyle.Render(h.Key)+" "+footerDescStyle.Render(h.Desc))
	}
	left := strings.Join(parts, "  ")

	right := onlineStyle.Render(fmt.Sprintf("● %d online", activeUsers))

	// Calculate padding between left and right
	leftWidth := ansi.StringWidth(left)
	rightWidth := ansi.StringWidth(right)
	innerWidth := width - 2 // account for padding
	gap := innerWidth - leftWidth - rightWidth
	if gap < 1 {
		gap = 1
	}

	bar := left + strings.Repeat(" ", gap) + right
	return footerBarStyle.Width(width).Render(bar)
}

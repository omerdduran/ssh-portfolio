package ui

import "charm.land/lipgloss/v2"

var (
	// Dracula palette
	purple    = lipgloss.Color("#bd93f9")
	pink      = lipgloss.Color("#ff79c6")
	cyan      = lipgloss.Color("#8be9fd")
	green     = lipgloss.Color("#50fa7b")
	orange    = lipgloss.Color("#ffb86c")
	red       = lipgloss.Color("#ff5555")
	yellow    = lipgloss.Color("#f1fa8c")
	gray      = lipgloss.Color("#6272a4")
	white     = lipgloss.Color("#f8f8f2")
	darkGray  = lipgloss.Color("#44475a")
	bg        = lipgloss.Color("#282a36")
	currentBg = lipgloss.Color("#44475a")

	// Text styles
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(purple).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(pink)

	infoStyle = lipgloss.NewStyle().
			Foreground(cyan)

	quitStyle = lipgloss.NewStyle().
			Foreground(gray).
			MarginTop(1)

	// Container
	containerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(purple).
			Padding(2, 4).
			Foreground(white)

	// List styles
	listItemStyle = lipgloss.NewStyle().
			Foreground(white).
			PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(green).
				Bold(true).
				PaddingLeft(2)

	listCountStyle = lipgloss.NewStyle().
			Foreground(gray)

	// Detail view
	detailTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(purple).
				MarginBottom(1)

	detailMetaStyle = lipgloss.NewStyle().
			Foreground(gray)

	// Tags
	tagStyle = lipgloss.NewStyle().
			Foreground(bg).
			Background(cyan).
			Padding(0, 1)

	// Links
	linkStyle = lipgloss.NewStyle().
			Foreground(cyan).
			Underline(true)

	// Timeline (work)
	companyStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(green)

	roleStyle = lipgloss.NewStyle().
			Foreground(pink)

	dateRangeStyle = lipgloss.NewStyle().
			Foreground(gray)

	timelineMarker = lipgloss.NewStyle().
			Foreground(purple).
			Bold(true)

	// Header & footer
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(purple).
			Background(currentBg).
			Padding(0, 2)

	footerStyle = lipgloss.NewStyle().
			Foreground(gray).
			Background(currentBg).
			Padding(0, 1)

	footerKeyStyle = lipgloss.NewStyle().
			Foreground(white).
			Bold(true)

	footerDescStyle = lipgloss.NewStyle().
			Foreground(gray)

	// Highlight
	highlightStyle = lipgloss.NewStyle().
			Foreground(yellow).
			Bold(true)

	// Separator
	separatorStyle = lipgloss.NewStyle().
			Foreground(darkGray)
)

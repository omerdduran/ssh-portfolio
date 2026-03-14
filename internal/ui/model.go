package ui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/x/ansi"
	"github.com/omerdduran/ssh-portfolio/internal/content"
	"github.com/omerdduran/ssh-portfolio/internal/ui/views"
)

type Page int

const (
	PageHome Page = iota
	PageMenu
	PageBlogList
	PageBlogDetail
	PageProjectsList
	PageProjectDetail
	PageWork
	PageChangelog
)

type Model struct {
	width  int
	height int
	user   string

	page         Page
	menuCursor   int
	listCursor   int
	scrollOffset int

	content *content.SiteContent
	rain    RainState
}

func TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	m := Model{
		width:   pty.Window.Width,
		height:  pty.Window.Height,
		user:    s.User(),
		page:    PageHome,
		content: content.Get(),
		rain:    newRainState(pty.Window.Width, pty.Window.Height),
	}
	return m, nil
}

func (m Model) Init() tea.Cmd {
	return rainTick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case RainTickMsg:
		if m.page == PageHome {
			m.rain.update()
			return m, rainTick()
		}
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.rain.resize(msg.Width, msg.Height)

	case tea.KeyPressMsg:
		code := msg.Code

		// Global quit
		if code == 'c' && msg.Mod == tea.ModCtrl {
			return m, tea.Quit
		}
		if code == 'q' && msg.Mod == 0 {
			return m, tea.Quit
		}

		// Back navigation (esc or backspace)
		isBack := code == tea.KeyEscape || code == tea.KeyBackspace

		switch m.page {
		case PageHome:
			if code == tea.KeyEnter {
				m.page = PageMenu
				m.menuCursor = 0
			}

		case PageMenu:
			switch {
			case code == 'k' || code == tea.KeyUp:
				if m.menuCursor > 0 {
					m.menuCursor--
				}
			case code == 'j' || code == tea.KeyDown:
				if m.menuCursor < 3 {
					m.menuCursor++
				}
			case code == tea.KeyEnter:
				m.listCursor = 0
				m.scrollOffset = 0
				switch m.menuCursor {
				case 0:
					m.page = PageBlogList
				case 1:
					m.page = PageProjectsList
				case 2:
					m.page = PageWork
				case 3:
					m.page = PageChangelog
				}
			case isBack:
				m.page = PageHome
				return m, rainTick()
			}

		case PageBlogList:
			maxIdx := len(m.content.Blog) - 1
			if maxIdx < 0 {
				maxIdx = 0
			}
			switch {
			case code == 'k' || code == tea.KeyUp:
				if m.listCursor > 0 {
					m.listCursor--
				}
			case code == 'j' || code == tea.KeyDown:
				if m.listCursor < maxIdx {
					m.listCursor++
				}
			case code == tea.KeyEnter:
				if len(m.content.Blog) > 0 {
					m.scrollOffset = 0
					m.page = PageBlogDetail
				}
			case isBack:
				m.page = PageMenu
				m.listCursor = 0
			}

		case PageProjectsList:
			maxIdx := len(m.content.Projects) - 1
			if maxIdx < 0 {
				maxIdx = 0
			}
			switch {
			case code == 'k' || code == tea.KeyUp:
				if m.listCursor > 0 {
					m.listCursor--
				}
			case code == 'j' || code == tea.KeyDown:
				if m.listCursor < maxIdx {
					m.listCursor++
				}
			case code == tea.KeyEnter:
				if len(m.content.Projects) > 0 {
					m.scrollOffset = 0
					m.page = PageProjectDetail
				}
			case isBack:
				m.page = PageMenu
				m.listCursor = 0
			}

		case PageBlogDetail, PageProjectDetail, PageWork, PageChangelog:
			switch {
			case code == 'k' || code == tea.KeyUp:
				if m.scrollOffset > 0 {
					m.scrollOffset--
				}
			case code == 'j' || code == tea.KeyDown:
				m.scrollOffset++
			case code == 'g':
				m.scrollOffset = 0
			case code == 'G':
				m.scrollOffset = 9999
			case isBack:
				m.scrollOffset = 0
				switch m.page {
				case PageBlogDetail:
					m.page = PageBlogList
				case PageProjectDetail:
					m.page = PageProjectsList
				case PageWork, PageChangelog:
					m.page = PageMenu
				}
			}
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	var screen string

	switch m.page {
	case PageHome:
		screen = m.homeWithRain()

	case PageMenu:
		items := []views.MenuItem{
			{Title: "Blog", Desc: "Read my thoughts and writings", Count: len(m.content.Blog)},
			{Title: "Projects", Desc: "Software I've built", Count: len(m.content.Projects)},
			{Title: "Work", Desc: "Professional experience", Count: len(m.content.Work)},
			{Title: "Changelog", Desc: "What's new on the site", Count: len(m.content.Changelog)},
		}
		screen = views.MenuView(items, m.menuCursor, m.width, m.height)

	case PageBlogList:
		screen = views.BlogListView(m.content.Blog, m.listCursor, m.width, m.height)

	case PageBlogDetail:
		if m.listCursor < len(m.content.Blog) {
			screen = views.BlogDetailView(m.content.Blog[m.listCursor], m.scrollOffset, m.width, m.height)
		}

	case PageProjectsList:
		screen = views.ProjectsListView(m.content.Projects, m.listCursor, m.width, m.height)

	case PageProjectDetail:
		if m.listCursor < len(m.content.Projects) {
			screen = views.ProjectDetailView(m.content.Projects[m.listCursor], m.scrollOffset, m.width, m.height)
		}

	case PageWork:
		screen = views.WorkView(m.content.Work, m.scrollOffset, m.width, m.height)

	case PageChangelog:
		screen = views.ChangelogView(m.content.Changelog, m.scrollOffset, m.width, m.height)
	}

	v := tea.NewView(screen)
	v.AltScreen = true
	return v
}

func (m Model) homeWithRain() string {
	if m.width <= 0 || m.height <= 0 {
		return ""
	}

	// Get the card (uncentered)
	card := views.HomeCard(m.user, m.width)
	cardLines := strings.Split(card, "\n")

	// Measure card dimensions using display width
	cardH := len(cardLines)
	cardW := 0
	for _, line := range cardLines {
		w := ansi.StringWidth(line)
		if w > cardW {
			cardW = w
		}
	}

	// Center position
	cardLeft := (m.width - cardW) / 2
	cardTop := (m.height - cardH) / 2
	if cardLeft < 0 {
		cardLeft = 0
	}
	if cardTop < 0 {
		cardTop = 0
	}

	// Get rain grid
	rainGrid := m.rain.renderGrid()

	// Composite rain + card
	lines := make([]string, m.height)
	for r := 0; r < m.height; r++ {
		if r < cardTop || r >= cardTop+cardH {
			// Full rain line
			if r < len(rainGrid) {
				lines[r] = renderRainLine(rainGrid[r])
			} else {
				lines[r] = strings.Repeat(" ", m.width)
			}
		} else {
			// Card row — composite: rain | card | rain
			cardRow := r - cardTop
			cardLine := ""
			if cardRow < len(cardLines) {
				cardLine = cardLines[cardRow]
			}

			var sb strings.Builder

			// Left rain margin
			if r < len(rainGrid) {
				for c := 0; c < cardLeft && c < len(rainGrid[r]); c++ {
					sb.WriteString(rainGrid[r][c])
				}
			}

			// Card content
			sb.WriteString(cardLine)

			// Right rain margin
			cardRight := cardLeft + cardW
			if r < len(rainGrid) {
				for c := cardRight; c < m.width && c < len(rainGrid[r]); c++ {
					sb.WriteString(rainGrid[r][c])
				}
			}

			lines[r] = sb.String()
		}
	}

	return strings.Join(lines, "\n")
}

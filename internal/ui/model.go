package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/ssh"
)

type Model struct {
	width  int
	height int
	user   string
}

func TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	m := Model{
		width:  pty.Window.Width,
		height: pty.Window.Height,
		user:   s.User(),
	}
	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() tea.View {
	title := titleStyle.Render("Ömer Duran")
	subtitle := subtitleStyle.Render("Software Engineer")
	info := infoStyle.Render("Welcome to my SSH portfolio!")

	greeting := ""
	if m.user != "" {
		greeting = infoStyle.Render(fmt.Sprintf("Hello, %s!", m.user))
	}

	quit := quitStyle.Render("Press q to exit")

	content := lipgloss.JoinVertical(lipgloss.Center,
		title,
		subtitle,
		"",
		info,
		greeting,
		quit,
	)

	box := containerStyle.Render(content)

	screen := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
	v := tea.NewView(screen)
	v.AltScreen = true
	return v
}

package views

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

var (
	homeTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#bd93f9"))

	homeSubtitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ff79c6"))

	homeInfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f8f8f2"))

	homeGreetingStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#8be9fd"))

	homeSocialStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8be9fd"))

	homeHintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6272a4"))

	homeSepStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#bd93f9"))

	homeContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#bd93f9")).
				Foreground(lipgloss.Color("#f8f8f2"))

	homeNameBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("#bd93f9")).
				Padding(0, 2).
				Bold(true).
				Foreground(lipgloss.Color("#bd93f9"))
)

func spacedName(name string) string {
	runes := []rune(strings.ToUpper(name))
	var parts []string
	for i, r := range runes {
		if r == ' ' {
			parts = append(parts, "   ")
		} else {
			parts = append(parts, string(r))
			if i < len(runes)-1 && runes[i+1] != ' ' {
				parts = append(parts, " ")
			}
		}
	}
	return strings.Join(parts, "")
}

func HomeView(user string, width, height int) string {
	narrow := width < 60

	// Name banner
	spaced := spacedName("Ömer Duran")
	var nameBlock string
	if narrow {
		nameBlock = homeTitleStyle.Render(spaced)
	} else {
		nameBlock = homeNameBoxStyle.Render(spaced)
	}

	// Subtitle
	subtitle := homeSubtitleStyle.Render("Software Developer")

	// Separator
	sepWidth := 35
	if narrow {
		sepWidth = 25
	}
	sep := homeSepStyle.Render(strings.Repeat("━", sepWidth))

	// Welcome paragraph
	welcome := homeInfoStyle.Render(
		"Welcome to my terminal portfolio.\n" +
			"Explore my blog, projects, and\n" +
			"work experience — all from your\n" +
			"terminal.",
	)

	// Greeting
	greeting := ""
	if user != "" {
		greeting = homeGreetingStyle.Render(fmt.Sprintf("Hello, %s!", user))
	}

	// Social links (clickable OSC 8 hyperlinks)
	socialLinkStyle := homeSocialStyle.Underline(true)
	socialDot := homeSocialStyle.Render(" · ")
	socials := strings.Join([]string{
		socialLinkStyle.Hyperlink("https://github.com/omerdduran").Render("github"),
		socialDot,
		socialLinkStyle.Hyperlink("https://www.linkedin.com/in/omerdduran").Render("linkedin"),
		socialDot,
		socialLinkStyle.Hyperlink("https://mastodon.social/@omerdduran").Render("mastodon"),
		socialDot,
		socialLinkStyle.Hyperlink("https://www.omerduran.dev").Render("web"),
	}, "")

	// Hint
	hint := homeHintStyle.Render("↵ Enter to explore  ·  q to quit")

	// Build content
	parts := []string{
		"",
		nameBlock,
		subtitle,
		"",
		sep,
		"",
		welcome,
		"",
	}
	if greeting != "" {
		parts = append(parts, greeting, "")
	}
	parts = append(parts,
		socials,
		"",
		sep,
		"",
		hint,
		"",
	)

	content := lipgloss.JoinVertical(lipgloss.Center, parts...)

	padding := 4
	if narrow {
		padding = 2
	}
	box := homeContainerStyle.Padding(2, padding).Render(content)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}

// HomeCard returns the home card without centering (for compositing with rain).
func HomeCard(user string, width int) string {
	narrow := width < 60

	spaced := spacedName("Ömer Duran")
	var nameBlock string
	if narrow {
		nameBlock = homeTitleStyle.Render(spaced)
	} else {
		nameBlock = homeNameBoxStyle.Render(spaced)
	}

	subtitle := homeSubtitleStyle.Render("Software Developer")

	sepWidth := 35
	if narrow {
		sepWidth = 25
	}
	sep := homeSepStyle.Render(strings.Repeat("━", sepWidth))

	welcome := homeInfoStyle.Render(
		"Welcome to my terminal portfolio.\n" +
			"Explore my blog, projects, and\n" +
			"work experience — all from your\n" +
			"terminal.",
	)

	greeting := ""
	if user != "" {
		greeting = homeGreetingStyle.Render(fmt.Sprintf("Hello, %s!", user))
	}

	socialLinkStyle := homeSocialStyle.Underline(true)
	socialDot := homeSocialStyle.Render(" · ")
	socials := strings.Join([]string{
		socialLinkStyle.Hyperlink("https://github.com/omerdduran").Render("github"),
		socialDot,
		socialLinkStyle.Hyperlink("https://www.linkedin.com/in/omerdduran").Render("linkedin"),
		socialDot,
		socialLinkStyle.Hyperlink("https://mastodon.social/@omerdduran").Render("mastodon"),
		socialDot,
		socialLinkStyle.Hyperlink("https://www.omerduran.dev").Render("web"),
	}, "")

	hint := homeHintStyle.Render("↵ Enter to explore  ·  q to quit")

	parts := []string{
		"",
		nameBlock,
		subtitle,
		"",
		sep,
		"",
		welcome,
		"",
	}
	if greeting != "" {
		parts = append(parts, greeting, "")
	}
	parts = append(parts,
		socials,
		"",
		sep,
		"",
		hint,
		"",
	)

	content := lipgloss.JoinVertical(lipgloss.Center, parts...)

	padding := 4
	if narrow {
		padding = 2
	}
	return homeContainerStyle.Padding(2, padding).Render(content)
}

package components

import (
	"charm.land/glamour/v2"
	"charm.land/glamour/v2/styles"
)

func RenderMarkdown(source string, width int) string {
	if width < 20 {
		width = 20
	}
	if width > 120 {
		width = 120
	}

	r, err := glamour.NewTermRenderer(
		glamour.WithStandardStyle(styles.DraculaStyle),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return source
	}

	out, err := r.Render(source)
	if err != nil {
		return source
	}
	return out
}

package ui

import (
	"math/rand"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// Half-width Katakana + digits
var rainChars = []rune(
	"ｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜｦﾝ0123456789",
)

// Pre-computed styles: index 0 = head (bright green), 1..8 = trail (dimming)
var rainStyles [9]lipgloss.Style

func init() {
	greens := [9]string{
		"#50fa7b", // head
		"#45d96c", "#3ab85d", "#2f974e", "#24763f",
		"#195530", "#0e3421", "#071a11", "#030d08",
	}
	for i, c := range greens {
		rainStyles[i] = lipgloss.NewStyle().Foreground(lipgloss.Color(c))
	}
}

type RainDrop struct {
	col      int
	row      float64
	speed    float64
	trailLen int
	delay    int // ticks before visible
	chars    []rune
}

type RainState struct {
	drops         []RainDrop
	width, height int
	rng           *rand.Rand
	frame         int
}

type RainTickMsg time.Time

func rainTick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return RainTickMsg(t)
	})
}

func newRainState(w, h int) RainState {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rs := RainState{width: w, height: h, rng: rng}
	rs.generateDrops()
	return rs
}

func (rs *RainState) generateDrops() {
	if rs.width <= 0 || rs.height <= 0 {
		rs.drops = nil
		return
	}
	// ~1 drop per 4 columns
	count := rs.width / 4
	if count < 1 {
		count = 1
	}
	rs.drops = make([]RainDrop, count)
	for i := range rs.drops {
		rs.drops[i] = rs.newDrop(true)
	}
}

func (rs *RainState) newDrop(stagger bool) RainDrop {
	trailLen := 5 + rs.rng.Intn(11) // 5-15
	chars := make([]rune, trailLen+1)
	for j := range chars {
		chars[j] = rainChars[rs.rng.Intn(len(rainChars))]
	}
	delay := 0
	startRow := float64(0)
	if stagger {
		// Spread initial positions across the screen
		startRow = -float64(rs.rng.Intn(rs.height + 10))
		delay = rs.rng.Intn(20)
	}
	return RainDrop{
		col:      rs.rng.Intn(rs.width),
		row:      startRow,
		speed:    0.5 + rs.rng.Float64()*1.5, // 0.5-2.0
		trailLen: trailLen,
		delay:    delay,
		chars:    chars,
	}
}

func (rs *RainState) update() {
	rs.frame++
	for i := range rs.drops {
		d := &rs.drops[i]
		if d.delay > 0 {
			d.delay--
			continue
		}
		d.row += d.speed

		// Shimmer: cycle one random char per frame
		idx := rs.rng.Intn(len(d.chars))
		d.chars[idx] = rainChars[rs.rng.Intn(len(rainChars))]

		// Respawn when entire trail is off screen
		if int(d.row)-d.trailLen > rs.height {
			*d = rs.newDrop(false)
			d.row = -float64(rs.rng.Intn(5))
		}
	}
}

func (rs *RainState) resize(w, h int) {
	rs.width = w
	rs.height = h
	rs.generateDrops()
}

// renderGrid returns a height×width grid of styled single-char strings.
// Empty cells are " ".
func (rs *RainState) renderGrid() [][]string {
	grid := make([][]string, rs.height)
	for r := range grid {
		row := make([]string, rs.width)
		for c := range row {
			row[c] = " "
		}
		grid[r] = row
	}

	for i := range rs.drops {
		d := &rs.drops[i]
		if d.delay > 0 {
			continue
		}
		headRow := int(d.row)
		for t := 0; t <= d.trailLen; t++ {
			r := headRow - t
			if r < 0 || r >= rs.height {
				continue
			}
			if d.col < 0 || d.col >= rs.width {
				continue
			}
			charIdx := t % len(d.chars)
			styleIdx := t
			if styleIdx > 8 {
				styleIdx = 8
			}
			grid[r][d.col] = rainStyles[styleIdx].Render(string(d.chars[charIdx]))
		}
	}

	return grid
}

// renderLine joins a single row of the grid into a string.
func renderRainLine(row []string) string {
	return strings.Join(row, "")
}

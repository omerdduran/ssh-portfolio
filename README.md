# ssh-portfolio

A terminal portfolio served over SSH, featuring a Matrix rain animation, blog posts, projects, and work history — all rendered in your terminal.

Built with [Wish](https://github.com/charmbracelet/wish), [Bubble Tea](https://github.com/charmbracelet/bubbletea), and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

## Features

- Matrix-style rain animation on the home page
- Blog posts, projects, work experience, and changelog sections
- Live content fetched from [omerduran.dev](https://www.omerduran.dev) with offline fallback data
- Vim-style navigation (`j`/`k`, `g`/`G`)
- Clickable hyperlinks (in supported terminals)
- Personalized greeting using your SSH username
- Responsive layout adapting to terminal size

## Try it

```bash
ssh localhost -p 23234
```

## Run locally

```bash
go run .
```
## Docker

```bash
docker build -t ssh-portfolio .
docker run -p 23234:23234 ssh-portfolio
```

## Configuration

| Variable | Default | Description |
|---|---|---|
| `PORTFOLIO_URL` | `https://www.omerduran.dev` | Base URL for the content API |

## Navigation

| Key | Action |
|---|---|
| `Enter` | Select / open |
| `Esc` / `Backspace` | Go back |
| `j` / `Down` | Move down |
| `k` / `Up` | Move up |
| `g` | Scroll to top |
| `G` | Scroll to bottom |
| `q` / `Ctrl+C` | Quit |

## Project structure

```
.
├── main.go                  # SSH server setup
├── internal/
│   ├── content/             # API fetcher, cache, fallback data
│   └── ui/
│       ├── model.go         # Bubble Tea model & navigation
│       ├── rain.go          # Matrix rain animation
│       ├── styles.go        # Shared styles
│       ├── components/      # Header, footer, markdown renderer
│       └── views/           # Page views (home, menu, blog, projects, work, changelog)
└── Dockerfile
```

## License

[MIT](LICENSE)

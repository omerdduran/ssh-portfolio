# ssh-portfolio

A terminal-accessible portfolio served over SSH, built with [Wish](https://github.com/charmbracelet/wish), [Bubble Tea](https://github.com/charmbracelet/bubbletea), and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

## Usage

```bash
ssh localhost -p 23234
```

Press `q` to exit.

## Run locally

```bash
go run .
```

## Docker

```bash
docker build -t ssh-portfolio .
docker run -p 23234:23234 ssh-portfolio
```

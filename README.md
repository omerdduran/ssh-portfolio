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
├── Dockerfile
└── .woodpecker.yml          # Preview + release CI pipeline
```

## Deployment

Continuous deployment is driven by a Woodpecker pipeline (`.woodpecker.yml`) running
against a Forgejo mirror of this repo. Two long-lived branches produce two
independent environments:

| Branch    | Image tag                  | Host port        | Deploy dir                     |
|-----------|----------------------------|------------------|--------------------------------|
| `develop` | `ssh-portfolio:preview`    | `0.0.0.0:23235`  | `/opt/ssh-portfolio/preview`   |
| `release` | `ssh-portfolio:release`    | `0.0.0.0:23234`  | `/opt/ssh-portfolio/prod`      |

On each push, the pipeline builds the Docker image locally on the runner
(`docker:28-cli` with the host's Docker socket mounted) and then runs
`docker compose up -d` inside the matching deploy dir.

### Server layout

```
/opt/ssh-portfolio/
├── .ssh/                    # persisted Wish host key (owned by ssh-portfolio:ssh-portfolio, uid/gid 996:986)
│   ├── id_ed25519
│   └── id_ed25519.pub
├── preview/
│   └── docker-compose.yml
└── prod/
    └── docker-compose.yml
```

Both compose files mount `/opt/ssh-portfolio/.ssh` read-only into
`/app/.ssh` and set `user: "996:986"` so the non-root container can read the
host key. That way the SSH fingerprint stays stable across image rebuilds.

### Compose contract

```yaml
services:
  ssh:
    image: ssh-portfolio:release       # or :preview
    container_name: ssh-portfolio-prod # or ssh-portfolio-preview
    restart: unless-stopped
    user: "996:986"
    volumes:
      - /opt/ssh-portfolio/.ssh:/app/.ssh:ro
    ports:
      - "0.0.0.0:23234:23234"          # preview uses 23235
```

### Trust flag

The pipeline bind-mounts `/var/run/docker.sock` and `/opt/ssh-portfolio`, so
the Woodpecker repo must have the **volumes** trust flag enabled (Repo →
Settings → Trusted). Without it pipelines fail at the linter with
`Insufficient trust level to use volumes`.

### Host SSH key

The key under `/opt/ssh-portfolio/.ssh/` is the long-lived server identity
that clients pin. Do **not** regenerate it — losing it changes the server
fingerprint and breaks every existing `known_hosts` entry. If you ever wipe
the deploy dir, restore this directory from backup first.

### First-time bootstrap

```bash
# On the server (one-time)
useradd -r -s /usr/sbin/nologin ssh-portfolio       # uid/gid used by the containers
mkdir -p /opt/ssh-portfolio/{preview,prod,.ssh}
chown -R ssh-portfolio:ssh-portfolio /opt/ssh-portfolio/.ssh
chmod 700 /opt/ssh-portfolio/.ssh
ssh-keygen -t ed25519 -N "" -f /opt/ssh-portfolio/.ssh/id_ed25519
# drop the two docker-compose.yml files shown above into preview/ and prod/
```

Then push `develop` or `release` to trigger the first build.

## License

[MIT](LICENSE)

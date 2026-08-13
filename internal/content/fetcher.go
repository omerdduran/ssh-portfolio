package content

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

var baseURL = "https://www.omerduran.dev"

func init() {
	if u := os.Getenv("PORTFOLIO_URL"); u != "" {
		baseURL = u
	}
}

func FetchAll() (*SiteContent, error) {
	sc := &SiteContent{}
	g := new(errgroup.Group)

	g.Go(func() error {
		projects, err := fetchJSON[[]Project]("/api/projects.json")
		if err != nil {
			return fmt.Errorf("projects: %w", err)
		}
		sc.Projects = projects
		return nil
	})

	g.Go(func() error {
		work, err := fetchJSON[[]WorkEntry]("/api/work.json")
		if err != nil {
			return fmt.Errorf("work: %w", err)
		}
		sc.Work = work
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return sc, nil
}

func fetchJSON[T any](path string) (T, error) {
	var zero T
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(baseURL + path)
	if err != nil {
		return zero, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return zero, fmt.Errorf("HTTP %d from %s", resp.StatusCode, path)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return zero, err
	}

	var result T
	if err := json.Unmarshal(body, &result); err != nil {
		return zero, err
	}
	return result, nil
}

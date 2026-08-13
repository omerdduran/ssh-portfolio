package content

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var baseURL = "https://www.omerduran.dev"

func init() {
	if u := os.Getenv("PORTFOLIO_URL"); u != "" {
		baseURL = u
	}
}

// FetchAll loads each content type in parallel and returns whatever it managed
// to get. One endpoint failing no longer discards the rest: this ran under an
// errgroup before, so a single 404 aborted the whole sync and the app quietly
// served the embedded snapshot instead — which is exactly what happened for as
// long as /api/changelog.json was requested but never existed. An error comes
// back only when every source failed.
func FetchAll() (*SiteContent, error) {
	sc := &SiteContent{}

	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		errs []error
	)

	fetch := func(name string, load func() error) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := load(); err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("%s: %w", name, err))
				mu.Unlock()
			}
		}()
	}

	const sources = 2

	fetch("projects", func() error {
		projects, err := fetchJSON[[]Project]("/api/projects.json")
		if err != nil {
			return err
		}
		sc.Projects = projects
		return nil
	})

	fetch("work", func() error {
		work, err := fetchJSON[[]WorkEntry]("/api/work.json")
		if err != nil {
			return err
		}
		sc.Work = work
		return nil
	})

	wg.Wait()

	if len(errs) == sources {
		return nil, errors.Join(errs...)
	}
	if len(errs) > 0 {
		log.Printf("Content fetch partially failed: %v", errors.Join(errs...))
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

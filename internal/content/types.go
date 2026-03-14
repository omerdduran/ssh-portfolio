package content

import "time"

type BlogPost struct {
	Slug        string   `json:"slug"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Date        time.Time `json:"date"`
	Tags        []string `json:"tags"`
	Body        string   `json:"body"`
}

type Project struct {
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	DemoURL     string    `json:"demoURL"`
	RepoURL     string    `json:"repoURL"`
	Tags        []string  `json:"tags"`
	Body        string    `json:"body"`
}

type WorkEntry struct {
	Slug       string `json:"slug"`
	Company    string `json:"company"`
	Role       string `json:"role"`
	DateStart  string `json:"dateStart"`
	DateEnd    string `json:"dateEnd"`
	CompanyURL string `json:"companyUrl"`
	Body       string `json:"body"`
}

type ChangelogEntry struct {
	Slug      string   `json:"slug"`
	Title     string   `json:"title"`
	Date      string   `json:"date"`
	Summary   string   `json:"summary"`
	Tags      []string `json:"tags"`
	Highlight bool     `json:"highlight"`
	Body      string   `json:"body"`
}

type SiteContent struct {
	Blog      []BlogPost       `json:"blog"`
	Projects  []Project        `json:"projects"`
	Work      []WorkEntry      `json:"work"`
	Changelog []ChangelogEntry `json:"changelog"`
}

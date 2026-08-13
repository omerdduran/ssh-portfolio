package content

import "time"

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

type SiteContent struct {
	Projects []Project   `json:"projects"`
	Work     []WorkEntry `json:"work"`
}

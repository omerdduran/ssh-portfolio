package content

import (
	"embed"
	"encoding/json"
	"log"
)

//go:embed fallback_data/*.json
var fallbackFS embed.FS

func LoadFallback() *SiteContent {
	sc := &SiteContent{}

	unmarshalFile("fallback_data/blog.json", &sc.Blog)
	unmarshalFile("fallback_data/projects.json", &sc.Projects)
	unmarshalFile("fallback_data/work.json", &sc.Work)
	unmarshalFile("fallback_data/changelog.json", &sc.Changelog)

	return sc
}

func unmarshalFile(name string, dst any) {
	data, err := fallbackFS.ReadFile(name)
	if err != nil {
		log.Printf("Fallback read %s: %v", name, err)
		return
	}
	if err := json.Unmarshal(data, dst); err != nil {
		log.Printf("Fallback parse %s: %v", name, err)
	}
}

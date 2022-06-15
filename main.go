package main

import (
	"flag"
	"log"

	"github.com/mszostok/issue-shooter/pkg/issue"
)

// Config holds app configuration
type Config struct {
	IssueMeta   issue.Metadata
	RepoBaseURL string
}

func main() {
	cfg := Config{}
	flag.StringVar(&cfg.IssueMeta.Title, "title", "", "Supply a title. Will prompt for one otherwise.")
	flag.StringVar(&cfg.IssueMeta.Type, "type", "", "Issue type. Will prompt for one otherwise.")
	flag.BoolVar(&cfg.IssueMeta.EditBody, "edit-body", false, "Edit default body.")

	// TODO: this can be list with predefined repository against which you can open an issue.
	flag.StringVar(&cfg.RepoBaseURL, "repo-url", issue.DefaultBaseURL, "Specify repository base URL.")
	flag.Parse()

	err := issue.ResolveWithSurvey(&cfg.IssueMeta)
	fatalOnErr(err)

	err = issue.Open(cfg.IssueMeta, cfg.RepoBaseURL)
	fatalOnErr(err)
}

func fatalOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

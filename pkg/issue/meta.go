package issue

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/MakeNowJust/heredoc/v2"
)

//type Type string // can be introduced for below consts

const (
	Feature     = "feature-request 🚀"
	Bug         = "bug 🐛"
	TechDebt    = "tech-debt 🙈"
	Enhancement = "enhancement"
)

type Metadata struct {
	Title    string
	Body     string
	Type     string
	EditBody bool
}

func ResolveWithSurvey(meta *Metadata) error {
	var qs []*survey.Question

	if meta.Type == "" {
		qs = append(qs, &survey.Question{
			Name: "type",
			Prompt: &survey.Select{
				Message: "Choose a template",
				Options: []string{Feature, Bug, TechDebt, Enhancement},
			},
			Validate: survey.Required,
		})
	}

	if meta.Title == "" {
		qs = append(qs, &survey.Question{
			Name: "title",
			Prompt: &survey.Input{
				Message: "Title:",
			},
			Validate: survey.Required,
		})
	}

	if !meta.EditBody {
		qs = append(qs, &survey.Question{
			Name: "EditBody",
			Prompt: &survey.Confirm{
				Message: "Do you want to edit body?",
				Default: false,
			},
		})
	}

	if err := survey.Ask(qs, meta); err != nil {
		return fmt.Errorf("while asking about issue metadata: %w", err)
	}

	meta.Body = RenderBody(meta.Type)

	if !meta.EditBody { // open with default body, user can edit it directly in browser
		return nil
	}

	bodyEditor := &survey.Editor{
		Message:       "Add extra information to issue body",
		Default:       meta.Body,
		AppendDefault: true,
		HideDefault:   true,
		FileName:      "*.md",
	}
	if err := survey.AskOne(bodyEditor, &meta.Body); err != nil {
		return fmt.Errorf("while opening editor for body edit: %w", err)
	}

	return nil
}

// RenderBody returns default body for a given issue type.
// TODO:
//   - template can be fetched directly from GitHub
func RenderBody(issue string) string {
	switch issue {
	case Bug:
		return heredoc.Doc(`
		  **Describe the bug**
		  A clear and concise description of what the bug is.

		  **To Reproduce**
		  Steps to reproduce the behavior:
		    1. Run '...'
		    2. Specify '...'
		    3. See error

		  **Expected behavior**
		  A clear and concise description of what you expected to happen.

		  **Version / Cluster**
		  - Testkube version - 1.1.4
		  - What Kubernetes cluster type - EC2 (t2.medium instance)
		  - K8s version - 1.22.1
		  - Cypress version - 9.5.4
		`)
	case Enhancement:
		return heredoc.Doc(`
		  **Enhancement**
		`)
	case TechDebt:
		return heredoc.Doc(`
		  **Tech Debt**
		`)
	case Feature:
		return heredoc.Doc(`
		  **Feature**
		`)
	default:
		return ""
	}
}

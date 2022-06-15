package issue

import (
	"fmt"
	"os"

	prShared "github.com/cli/cli/v2/pkg/cmd/pr/shared"
	"github.com/cli/cli/v2/utils"
	"github.com/skratchdot/open-golang/open"
)

const DefaultBaseURL = "https://github.com/kubeshop/testkube/issues/new"

func Open(opts Metadata, baseURL string) error {
	issue := prShared.IssueMetadataState{
		Type:   prShared.IssueMetadata,
		Body:   opts.Body,
		Title:  opts.Title,
		Labels: []string{opts.Type}, // TODO: this is meh..
	}

	openURL, err := prShared.WithPrAndIssueQueryParams(nil, nil, baseURL, issue)
	if err != nil {
		return err
	}

	if !utils.ValidURL(openURL) {
		return fmt.Errorf("cannot open in browser: maximum URL length exceeded")
	}

	fmt.Fprintf(os.Stderr, "\nOpening %s in your browser.\n", utils.DisplayURL(openURL))

	return open.Start(openURL)
}

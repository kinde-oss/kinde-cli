package release

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v28/github"
)

var Branch = "master"

func IsNeedingUpdate() {
	if Branch != "master" {
		client := github.NewClient(nil)
		rep, _, err := client.Repositories.GetLatestRelease(context.Background(), "kinde-oss", "kinde-cli")

		latest := *rep.TagName

		if err != nil {
			return
		}

		if strings.TrimPrefix(Branch, "v") != strings.TrimPrefix(latest, "v") {
			fmt.Printf("An update is available: %v", latest)
		}
	}
}

package release

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v28/github"
)

var (
	Version = ""
	Commit  = ""
	Date    = ""
)

func IsNeedingUpdate() {
	client := github.NewClient(nil)
	rep, _, err := client.Repositories.GetLatestRelease(context.Background(), "kinde-oss", "kinde-cli")

	if err != nil {
		return
	}

	latest := *rep.TagName

	if strings.TrimPrefix(Version, "v") != strings.TrimPrefix(latest, "v") {
		fmt.Printf("An update is available: %v", latest)
	}
}

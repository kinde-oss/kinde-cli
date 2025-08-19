package release

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/google/go-github/v28/github"
)

var (
	Version = ""
	Commit  = ""
	Date    = ""
)

func IsNeedingUpdate() {
	client := github.NewClient(nil)
	fmt.Printf("Checking for updates... \n")
	s := spinner.New(spinner.CharSets[rand.Intn(90)], 100*time.Millisecond)
	s.Start()

	rep, _, err := client.Repositories.GetLatestRelease(context.Background(), "kinde-oss", "kinde-cli")

	s.Stop()

	if err != nil {
		return
	}

	latest := *rep.TagName
	if strings.TrimPrefix(Version, "v") != strings.TrimPrefix(latest, "v") {
		fmt.Printf("An update is available: %v\n", latest)
	}
}

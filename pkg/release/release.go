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
	s := spinner.New(spinner.CharSets[rand.Intn(len(spinner.CharSets))], 100*time.Millisecond)
	s.Start()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rep, _, err := client.Repositories.GetLatestRelease(ctx, "kinde-oss", "kinde-cli")

	s.Stop()

	if err != nil {
		return
	}
	latest := *rep.TagName
	if strings.TrimPrefix(Version, "v") != strings.TrimPrefix(latest, "v") {
		fmt.Printf("An update is available: %v\n", latest)
	}
}

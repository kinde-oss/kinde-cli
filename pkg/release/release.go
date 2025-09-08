package release

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
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
	latest := rep.GetTagName()
	if latest == "" {
		fmt.Println("Latest release tag not found; cannot check for updates")
		return
	}

	// Parse versions using semver
	currentVersion, err := semver.NewVersion(strings.TrimPrefix(Version, "v"))
	if err != nil {
		// Handle development versions or invalid semver
		fmt.Printf("Current version '%s' is not a semantic version (development build)\n", Version)
		fmt.Printf("Latest release: %v\n", latest)
		return
	}

	latestVersion, err := semver.NewVersion(strings.TrimPrefix(latest, "v"))
	if err != nil {
		fmt.Printf("Error parsing latest version '%s': %v\n", latest, err)
		return
	}

	// Compare versions using semver
	if currentVersion.LessThan(latestVersion) {
		fmt.Printf("An update is available: %v (current: %v)\n", latest, Version)
	} else if currentVersion.Equal(latestVersion) {
		fmt.Printf("You are using the latest version: %v\n", Version)
	} else {
		// This case handles when current version is newer than latest (e.g., pre-release or development builds)
		fmt.Printf("You are using a newer version than the latest release: %v (latest: %v)\n", Version, latest)
	}
}

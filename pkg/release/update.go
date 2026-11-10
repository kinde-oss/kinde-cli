package release

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
)

const (
	latestReleaseURL = "https://api.github.com/repos/kinde-oss/kinde-cli/releases/latest"
	// checkInterval is how often the automatic check asks GitHub.
	checkInterval = 24 * time.Hour
	checkTimeout  = 3 * time.Second
)

// Update describes a release newer than the running one.
type Update struct {
	Current, Latest string
	URL             string
}

// Message is the notice shown to the user.
func (u *Update) Message() string {
	return fmt.Sprintf("A new version of kinde is available: %s → %s\n%s", u.Current, u.Latest, u.URL)
}

// LatestRelease is a release's version tag and web page.
type LatestRelease struct {
	Tag string `json:"tag_name"`
	URL string `json:"html_url"`
}

// Checker looks for newer releases, remembering the last answer in a state
// file so the automatic check asks GitHub at most once a day.
type Checker struct {
	StatePath string
	Now       func() time.Time
	Fetch     func(ctx context.Context) (LatestRelease, error)
}

// NewChecker returns a Checker that asks GitHub and keeps its state in
// statePath.
func NewChecker(statePath string) *Checker {
	return &Checker{StatePath: statePath, Now: time.Now, Fetch: fetchLatestRelease}
}

type checkState struct {
	CheckedAt time.Time     `json:"checked_at"`
	Latest    LatestRelease `json:"latest"`
}

// Check returns the newer release, or nil if this build is up to date, is a
// development build, or the check failed. Unless force is set, a result less
// than a day old is reused instead of asking GitHub again.
func (c *Checker) Check(ctx context.Context, force bool) (*Update, error) {
	current, err := semver.NewVersion(strings.TrimPrefix(Version, "v"))
	if err != nil {
		return nil, nil // a development build has nothing to compare
	}

	state := c.readState()
	if force || state == nil || c.Now().Sub(state.CheckedAt) >= checkInterval {
		ctx, cancel := context.WithTimeout(ctx, checkTimeout)
		defer cancel()
		latest, err := c.Fetch(ctx)
		if err != nil {
			// Record the attempt, keeping the last known release, so a
			// network problem delays one command a day rather than every
			// command.
			failed := &checkState{CheckedAt: c.Now()}
			if state != nil {
				failed.Latest = state.Latest
			}
			c.writeState(failed)
			return nil, err
		}
		state = &checkState{CheckedAt: c.Now(), Latest: latest}
		c.writeState(state)
	}
	if state.Latest.Tag == "" {
		return nil, nil // every check so far has failed
	}

	latest, err := semver.NewVersion(strings.TrimPrefix(state.Latest.Tag, "v"))
	if err != nil {
		return nil, fmt.Errorf("latest release has an invalid version %q", state.Latest.Tag)
	}
	if !current.LessThan(latest) {
		return nil, nil
	}
	return &Update{Current: current.String(), Latest: latest.String(), URL: state.Latest.URL}, nil
}

func (c *Checker) readState() *checkState {
	data, err := os.ReadFile(c.StatePath)
	if err != nil {
		return nil
	}
	var state checkState
	if json.Unmarshal(data, &state) != nil {
		return nil
	}
	return &state
}

// writeState saves the result; failing to is harmless, it only means the
// next run asks GitHub again.
func (c *Checker) writeState(state *checkState) {
	if data, err := json.Marshal(state); err == nil {
		_ = os.WriteFile(c.StatePath, data, 0o600)
	}
}

func fetchLatestRelease(ctx context.Context) (LatestRelease, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, latestReleaseURL, nil)
	if err != nil {
		return LatestRelease{}, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return LatestRelease{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return LatestRelease{}, fmt.Errorf("GitHub returned %s", resp.Status)
	}

	var latest LatestRelease
	if err := json.NewDecoder(resp.Body).Decode(&latest); err != nil {
		return LatestRelease{}, err
	}
	if latest.Tag == "" {
		return LatestRelease{}, errors.New("GitHub returned a release without a tag")
	}
	return latest, nil
}

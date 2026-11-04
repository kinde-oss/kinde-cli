package release

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeChecker returns a Checker whose clock and GitHub are controlled by the
// test, and a counter of how often it asked GitHub.
func fakeChecker(t *testing.T, latest string) (*Checker, *time.Time, *int) {
	t.Helper()
	now := time.Date(2026, 11, 4, 9, 0, 0, 0, time.UTC)
	fetches := 0
	checker := &Checker{
		StatePath: filepath.Join(t.TempDir(), "update-check.json"),
		Now:       func() time.Time { return now },
		Fetch: func(context.Context) (LatestRelease, error) {
			fetches++
			return LatestRelease{Tag: latest, URL: "https://github.com/kinde-oss/kinde-cli/releases/tag/" + latest}, nil
		},
	}
	return checker, &now, &fetches
}

func TestCheck_ReportsNewerRelease(t *testing.T) {
	setVersion(t, "1.0.0", "", "")
	checker, _, _ := fakeChecker(t, "v1.1.0")

	update, err := checker.Check(context.Background(), false)

	require.NoError(t, err)
	require.NotNil(t, update)
	assert.Equal(t, "A new version of kinde is available: 1.0.0 → 1.1.0\nhttps://github.com/kinde-oss/kinde-cli/releases/tag/v1.1.0", update.Message())
}

func TestCheck_NothingWhenUpToDateOrNewer(t *testing.T) {
	for _, current := range []string{"1.1.0", "1.2.0-beta.1"} {
		setVersion(t, current, "", "")
		checker, _, _ := fakeChecker(t, "v1.1.0")

		update, err := checker.Check(context.Background(), false)

		require.NoError(t, err)
		assert.Nil(t, update, current)
	}
}

func TestCheck_AsksGitHubAtMostOnceADay(t *testing.T) {
	setVersion(t, "1.0.0", "", "")
	checker, now, fetches := fakeChecker(t, "v1.1.0")

	_, _ = checker.Check(context.Background(), false)
	*now = now.Add(23 * time.Hour)
	update, _ := checker.Check(context.Background(), false)
	assert.Equal(t, 1, *fetches, "a result less than a day old is reused")
	assert.NotNil(t, update, "the saved result still reports the update")

	*now = now.Add(2 * time.Hour)
	_, _ = checker.Check(context.Background(), false)
	assert.Equal(t, 2, *fetches)

	_, _ = checker.Check(context.Background(), true)
	assert.Equal(t, 3, *fetches, "force always asks")
}

func TestCheck_DevelopmentBuildIsNeverChecked(t *testing.T) {
	setVersion(t, "dev", "", "")
	checker, _, fetches := fakeChecker(t, "v1.1.0")

	update, err := checker.Check(context.Background(), true)

	assert.NoError(t, err)
	assert.Nil(t, update)
	assert.Zero(t, *fetches)
}

func TestCheck_FailureIsReported(t *testing.T) {
	setVersion(t, "1.0.0", "", "")
	checker, _, _ := fakeChecker(t, "v1.1.0")
	checker.Fetch = func(context.Context) (LatestRelease, error) { return LatestRelease{}, errors.New("offline") }

	update, err := checker.Check(context.Background(), false)

	assert.EqualError(t, err, "offline")
	assert.Nil(t, update)
}

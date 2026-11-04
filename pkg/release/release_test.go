package release

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// setVersion sets the build variables for one test.
func setVersion(t *testing.T, version, commit, date string) {
	t.Helper()
	oldVersion, oldCommit, oldDate := Version, Commit, Date
	Version, Commit, Date = version, commit, date
	t.Cleanup(func() { Version, Commit, Date = oldVersion, oldCommit, oldDate })
}

func TestString(t *testing.T) {
	tests := []struct {
		version, commit, date string
		want                  string
	}{
		{"1.0.0", "7bb53d3a1b2c3d4e5f", "2026-11-04T10:20:30Z", "kinde version 1.0.0 (7bb53d3, 2026-11-04)"},
		{"v1.0.0", "", "", "kinde version 1.0.0"},
		// The defaults in cmd/kinde for a local build.
		{"dev", "none", "unknown", "kinde version dev"},
		{"", "", "", "kinde version dev"},
	}
	for _, tt := range tests {
		setVersion(t, tt.version, tt.commit, tt.date)
		assert.Equal(t, tt.want, String())
	}
}

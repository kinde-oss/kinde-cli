// Package release holds the CLI's version, set at build time, and checks
// GitHub for newer releases.
package release

import (
	"fmt"
	"strings"
)

// Set by cmd/kinde from values GoReleaser passes with -ldflags.
var (
	Version = ""
	Commit  = ""
	Date    = ""
)

// String describes this build, e.g. "kinde version 1.0.0 (7bb53d3, 2026-11-04)".
// Details a local build doesn't have are left out.
func String() string {
	version := strings.TrimPrefix(Version, "v")
	if version == "" {
		version = "dev"
	}

	var details []string
	if commit := Commit; commit != "" && commit != "none" {
		if len(commit) > 7 {
			commit = commit[:7]
		}
		details = append(details, commit)
	}
	if date := Date; date != "" && date != "unknown" {
		// GoReleaser passes an RFC 3339 timestamp; the day is enough.
		date, _, _ = strings.Cut(date, "T")
		details = append(details, date)
	}

	if len(details) == 0 {
		return fmt.Sprintf("kinde version %s", version)
	}
	return fmt.Sprintf("kinde version %s (%s)", version, strings.Join(details, ", "))
}

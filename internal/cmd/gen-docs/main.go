// Command gen-docs writes the Markdown command reference, one file per
// command, from the CLI's own command tree.
//
// Usage: gen-docs <output dir>
//
// It is run by go generate in pkg/cmd, after the API docs are generated.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kinde-oss/kinde-cli/pkg/cmd"
	"github.com/kinde-oss/kinde-cli/pkg/config"
	"github.com/kinde-oss/kinde-cli/pkg/iostreams"
	"github.com/spf13/cobra/doc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: gen-docs <output dir>")
		os.Exit(2)
	}
	if err := run(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, "gen-docs:", err)
		os.Exit(1)
	}
}

func run(dir string) error {
	root, err := cmd.NewRootCmd(iostreams.System(), &config.Config{})
	if err != nil {
		return err
	}
	root.InitDefaultCompletionCmd()
	// Keep the output identical between runs, so changes show up in diffs.
	root.DisableAutoGenTag = true

	// Remove the previous reference so deleted commands don't leave pages behind.
	old, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return err
	}
	for _, file := range old {
		if err := os.Remove(file); err != nil {
			return err
		}
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	return doc.GenMarkdownTree(root, dir)
}

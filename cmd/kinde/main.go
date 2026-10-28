package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kinde-oss/kinde-cli/pkg/cmd"
	"github.com/kinde-oss/kinde-cli/pkg/release"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	os.Exit(run())
}

// run holds the program logic so its deferred calls complete before os.Exit,
// which would otherwise skip them.
func run() int {

	release.Version = version
	release.Commit = commit
	release.Date = date

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := cmd.Execute(ctx, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}
	return 0
}

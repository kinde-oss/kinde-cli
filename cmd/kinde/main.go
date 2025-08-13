package main

import (
	"context"
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

	release.Version = version
	release.Commit = version
	release.Date = version

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd.Execute(ctx)
}

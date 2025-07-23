package main

import (
	"context"
	"time"

	"github.com/kinde-oss/kinde-cli/pkg/cmd"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd.Execute(ctx)
}

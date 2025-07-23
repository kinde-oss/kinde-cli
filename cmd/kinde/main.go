package main

import (
	"context"

	"github.com/kinde-oss/kinde-cli/pkg/cmd"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	cmd.Execute(ctx)
}

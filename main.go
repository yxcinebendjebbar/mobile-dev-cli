package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"
	"mobile-dev/cmd"
)

func main() {
	if err := fang.Execute(context.TODO(), cmd.RootCmd); err != nil {
		os.Exit(1)
	}
}

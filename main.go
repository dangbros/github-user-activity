package main

import (
	"os"

	"github.com/dangbros/github-user-activity/internal/app"
)

func main() {
	if err := app.Run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		os.Exit(1)
	}
}

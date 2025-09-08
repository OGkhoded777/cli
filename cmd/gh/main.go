package main

import (
	"os"

	"github.com/cli/cli/v2/internal/ghcmd"
)

// Main function to start the CLI
func main() {
	code := ghcmd.Main()
	os.Exit(int(code))
}


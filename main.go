package main

import (
	"os"

	"github.com/allen0039/komari_agent/cmd"
)

func main() {
	cmd.Execute()
	os.Exit(0)
}

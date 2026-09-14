package main

import (
	"os"

	"github.com/cloudpunks/dashboardctl/pkg/command"
	"github.com/joho/godotenv"
)

func main() {
	if env := os.Getenv("DASHBOARDCTL_ENV_FILE"); env != "" {
		_ = godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}

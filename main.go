package main

import (
	"fmt"
	"os"

	"github.com/Luzifer/rconfig"
)

type Config struct {
	PrintVersion bool `flag:"version"`
}

var (
	ProjectVersion string
	ProjectBuild   string
)

func main() {
	c := &Config{}
	rconfig.Parse(c)

	if c.PrintVersion {
		fmt.Printf("go-versionining version %s build %s\n", ProjectVersion, ProjectBuild)
		os.Exit(0)
	}

	fmt.Println("Hello again")
}

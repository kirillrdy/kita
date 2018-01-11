package kita

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const environmentsPath = KitaBasePath + "envs/"

type Environment struct {
	requirements []string
}

func (environment *Environment) Require(path string) {
	environment.requirements = append(environment.requirements, path)
}

func (environment Environment) Execute() {

	var binPaths []string

	for _, requiredPath := range environment.requirements {
		versions := versionsThatContains(requiredPath)
		packageVersion, err := LatestVersion(versions)
		if err != nil {
			log.Panicf("Failed to get latest version for env: %v", environment)
		}
		binPaths = append(binPaths, packageVersion.BinPath())
	}

	cmd := exec.Command("/bin/bash")
	path := fmt.Sprintf("PATH=%s:%s", strings.Join(binPaths, ":"), os.Getenv("PATH"))
	currentEnv := os.Environ()
	cmd.Env = append(currentEnv, path)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	Crash(err)
}

package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

const jiraUrl = "https://certinia.atlassian.net"
const defaultProject = "APLAT"
const browserEnvar = "BROWSER"

func browserCmd() ([]string, error) {
	cmd := os.Getenv(browserEnvar)
	if cmd == "" {
		return nil, errors.New(browserEnvar + " environment variable empty or not set\n")
	}
	parts := strings.Fields(cmd)
	return parts, nil
}

func url(id string) string {
	prefix := jiraUrl + "/"
	if id[0] >= '0' && id[0] <= '9' {
		return prefix + defaultProject + "-" + id
	}
	return prefix + id
}

func main() {
	browserCmd, err := browserCmd()
	if err != nil {
		panic(err)
	}
	// TODO: Read from stdin (process both or exclusive?
	args := os.Args[1:]
	for _, arg := range args {
		fullCmd := append(browserCmd, url(arg))
		err := exec.Command(fullCmd[0], fullCmd[1:]...).Run()
		if err != nil {
			panic(err)
		}
	}
}

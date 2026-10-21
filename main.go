package main

import (
	"bufio"
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
		return nil, errors.New(browserEnvar + " environment variable empty or not set")
	}
	parts := strings.Fields(cmd)
	return parts, nil
}

func url(id string) string {
	prefix := jiraUrl + "/browse/"
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
	fileInfo, _ := os.Stdin.Stat()
	var tickets []string
	if (fileInfo.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ticket := strings.TrimSpace(scanner.Text())
			if ticket != "" {
				tickets = append(tickets, ticket)
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	args := os.Args[1:]
	for _, arg := range args {
		tickets = append(tickets, arg)
	}
	for _, ticket := range tickets {
		fullCmd := append(browserCmd, url(ticket))
		err := exec.Command(fullCmd[0], fullCmd[1:]...).Run()
		if err != nil {
			panic(err)
		}
	}
}

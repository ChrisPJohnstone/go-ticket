package main

import (
	"fmt"
	"os"
)

const jiraUrl = "https://certinia.atlassian.net"
const defaultProject = "APLAT"
const browserEnvar = "BROWSER"

func url(id string) string {
	prefix := jiraUrl + "/"
	if id[0] >= '0' && id[0] <= '9' {
		return prefix + defaultProject + "-" + id
	}
	return prefix + id
}

func main() {
	browser := os.Getenv(browserEnvar)
	if browser == "" {
		fmt.Printf("%s environment variable empty or not set\n", browserEnvar)
		panic(browser)
	}
	// TODO: Read from stdin (process both or exclusive?
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Printf("%s\n", url(arg))
	}
}

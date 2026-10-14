package main

import (
	"fmt"
	"os"
)

const jiraUrl = "https://certinia.atlassian.net"
const defaultProject = "APLAT"
const browserEnvar = "BROWSER"

func main() {
	browser := os.Getenv(browserEnvar)
	if browser == "" {
		fmt.Printf("%s environment variable empty or not set\n", browserEnvar)
		panic(browser)
	}
	fmt.Printf("BROWSER = '%s'\n", browser)
	// TODO: Read from stdin (process both or exclusive?
	args := os.Args[1:]
	for _, arg := range args {
		fmt.Printf("%s\n", arg)
	}
}

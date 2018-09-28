package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {
	var numberOfReleases int
	var owner string
	var repository string

	flag.IntVar(&numberOfReleases, "n", 1, "Number of releases.")
	flag.StringVar(&owner, "o", "balabit", "Owner.")
	flag.StringVar(&repository, "r", "syslog-ng", "Repository.")
	flag.Parse()

	client := github.NewClient(nil)
	releases, _, _ := client.Repositories.ListReleases(context.Background(), owner, repository, nil)

	for i := 0; i < len(releases) && i < numberOfReleases; i++ {
		release := releases[i]
		for _, asset := range release.Assets {
			fmt.Printf("%s:%d\n", *asset.Name, *asset.DownloadCount)
		}
	}
}

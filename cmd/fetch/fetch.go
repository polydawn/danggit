package main

import (
	"fmt"
	"os"

	libgit "github.com/libgit2/git2go"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <remote>\n", os.Args[0])
		os.Exit(1)
	}
	repo, err := libgit.OpenRepository(".")
	if err != nil {
		fmt.Printf("not a repo: %s\n", err)
		os.Exit(5)
	}
	remotes := repo.Remotes
	remoteURL := os.Args[1]
	fmt.Printf("Remote: %s \n", remoteURL)

	remote, err := remotes.CreateAnonymous(remoteURL)
	if err != nil {
		fmt.Printf("an ineffable miracle beyond description: %s\n", err)
		os.Exit(6)
	}

	// err = remote.ConnectFetch(nil, nil, nil)
	err = remote.Fetch(nil, nil, "")
	if err != nil {
		fmt.Printf("well aren't you a fetching creature: %s\n", err)
		os.Exit(7)
	}
}

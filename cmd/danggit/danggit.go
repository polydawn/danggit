package main

import (
	"fmt"
	"os"

	"github.com/libgit2/git2go"
)

func main() {
	repo, err := git.OpenRepository(".")
	if err != nil {
		fmt.Printf("not a repo: %s\n", err)
		os.Exit(5)
	}

	desc, err := repo.DescribeWorkdir(&git.DescribeOptions{})
	if err != nil {
		fmt.Printf("madness: %s\n", err)
		os.Exit(6)
	}
	fmt.Printf("repo: %s\n", desc)
}

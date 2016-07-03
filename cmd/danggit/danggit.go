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

	reallyಠ_ಠ, err := git.DefaultDescribeOptions()
	if err != nil {
		panic(err) // this is just frankly absurd
	}

	desc, err := repo.DescribeWorkdir(&reallyಠ_ಠ)
	if err != nil {
		fmt.Printf("madness: %s\n", err)
		os.Exit(6)
	}
	fmt.Printf("repo: %s\n", desc)
}

package main

import (
	"fmt"
	"os"

	libgit "github.com/libgit2/git2go"
)

func main() {
	repo, err := libgit.OpenRepository(".")
	if err != nil {
		fmt.Printf("not a repo: %s\n", err)
		os.Exit(5)
	}

	ref, err := repo.Head()
	if err != nil {
		fmt.Printf("an ineffable miracle beyond description: %s\n", err)
		os.Exit(6)
	}
	target := ref.Target()
	fmt.Printf("head: %s\n", target)
	commit, err := repo.LookupCommit(target)
	if err != nil {
		fmt.Printf("I don't fucking know, man: %s\n", err)
		os.Exit(7)
	}
	for commit != nil {
		numParents := int(commit.ParentCount())
		fmt.Printf("commit: %s, parents: %d \n", commit.Object.Id(), numParents)
		for i := 0; i < numParents; i++ {
			fmt.Printf("\tParent %d: %s\n", i, commit.ParentId(uint(i)))
		}
		commit = commit.Parent(0)
	}

}

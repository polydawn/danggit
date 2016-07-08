package gat

import (
	libgit "github.com/libgit2/git2go"

	"polydawn.net/danggit/api"
)

func openRepository(pth git.LocalRepoPath) (*libgit.Repository, error) {
	return libgit.OpenRepositoryExtended(
		string(pth),
		libgit.RepositoryOpenNoSearch|libgit.RepositoryOpenBare,
		"",
	)
}

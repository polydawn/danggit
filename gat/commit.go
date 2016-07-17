package gat

import (
	libgit "github.com/libgit2/git2go"

	"polydawn.net/danggit/api"
)

type treeEntry struct {
	Filename string          // use full path
	Filemode libgit.Filemode // TODO git off my lawn c consts good lord
	Content  []byte          // for files
	// gitlinks not yet described by this struct
}

func createCommit(
	repo *libgit.Repository,
	commitHeaders *git.Commit, // treeID will be replaced
	contents []treeEntry,
) git.CommitID {
	if commitHeaders.TreeID != "" {
		panic("misuse: treeID will be set based on the contents you provided")
	}

	// assemble tree
	var treeOid *libgit.Oid
	func() {
		treeBuilder, err := repo.TreeBuilder()
		maybePanic(err)
		defer treeBuilder.Free()

		for _, content := range contents {
			switch content.Filemode {
			case libgit.FilemodeTree:
				panic("TODO")
			case libgit.FilemodeBlob, libgit.FilemodeBlobExecutable:
				// save blob
				blobOid, err := repo.CreateBlobFromBuffer(content.Content)
				maybePanic(err)
				// insert
				treeBuilder.Insert(content.Filename, blobOid, content.Filemode)
			case libgit.FilemodeLink:
				panic("TODO")
			case libgit.FilemodeCommit:
				panic("TODO")
			}

		}

		// write tree
		treeOid, err = treeBuilder.Write()
		maybePanic(err)
	}()
	// immediately look it up again because silly api
	tree, err := repo.LookupTree(treeOid)
	maybePanic(err)

	// assemble commit (mostly translating structs)
	commitOid, err := repo.CreateCommit(
		"",
		&libgit.Signature{
			Name:  commitHeaders.Author.Name,
			Email: commitHeaders.Author.Email,
			When:  commitHeaders.Author.When,
		},
		&libgit.Signature{
			Name:  commitHeaders.Committer.Name,
			Email: commitHeaders.Committer.Email,
			When:  commitHeaders.Committer.When,
		},
		commitHeaders.Message,
		tree,
	)
	maybePanic(err)
	return git.CommitID(commitOid.String())
}

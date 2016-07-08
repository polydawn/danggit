package gat

import (
	"fmt"
	"testing"
	"time"

	libgit "github.com/libgit2/git2go"
	. "github.com/smartystreets/goconvey/convey"

	"polydawn.net/danggit/lib/testutil"
)

func TestFetch(t *testing.T) {
	Convey("Given a local git repo", t, testutil.WithTmpdir(func() {
		repo, err := libgit.InitRepository("repo", true)
		maybePanic(err)

		Convey("and given some commits and branches", func() {
			// shove in file content
			blobOid, err := repo.CreateBlobFromBuffer([]byte("hello, world!\n"))
			maybePanic(err)

			// assemble tree
			treeBuilder, err := repo.TreeBuilder()
			maybePanic(err)
			defer treeBuilder.Free()
			treeBuilder.Insert("thefile", blobOid, libgit.FilemodeBlob)
			treeOid, err := treeBuilder.Write()
			maybePanic(err)
			// immediately look it up again because silly api
			tree, err := repo.LookupTree(treeOid)
			maybePanic(err)

			// assemble commit
			author := &libgit.Signature{
				Name:  "author",
				Email: "email@domain.wow",
				When:  time.Date(2009, 10, 14, 12, 0, 0, 0, time.UTC),
			}
			commitOid, err := repo.CreateCommit(
				"",
				author,
				author,
				"log message",
				tree,
			)
			maybePanic(err)
			// immediately look it up again because silly api
			commit, err := repo.LookupCommit(commitOid)
			maybePanic(err)

			// create branch
			branchName := "test-branch"
			branch, err := repo.CreateBranch(branchName, commit, false)
			maybePanic(err)
			branch.Free()

			remote, err := repo.Remotes.CreateAnonymous(repo.Path())
			maybePanic(err)

			remoteName := "test-remote"
			refspec := fmt.Sprintf("+refs/heads/*:refs/remotes/%s/*", remoteName)
			err = remote.Fetch([]string{refspec}, nil, "")
			maybePanic(err)

			expectedBranch := fmt.Sprintf("%s/%s", remoteName, branchName)
			remoteBranch, err := repo.LookupBranch(expectedBranch, libgit.BranchRemote)
			maybePanic(err)
			Convey("remote branch commit should be remote commit", func() {
				So(remoteBranch.Reference.Target(), ShouldResemble, commitOid)
			})

			Convey("execgit believe our work", func() {
				So(
					execgit.Bake("ls-remote", "repo").Output(),
					ShouldResemble,
					fmt.Sprintf("%s\trefs/heads/%s\n%s\trefs/remotes/%s\n", commitOid, branchName, commitOid, expectedBranch),
				)
			})
		})
	}))
}

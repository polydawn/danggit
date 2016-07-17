package gat

import (
	"testing"
	"time"

	libgit "github.com/libgit2/git2go"
	"github.com/polydawn/gosh"
	. "github.com/smartystreets/goconvey/convey"

	"polydawn.net/danggit/api"
	"polydawn.net/danggit/lib/testutil"
)

func TestCommit(t *testing.T) {
	Convey("Given a local git repo", t, testutil.WithTmpdir(func() {
		repo, err := libgit.InitRepository("repo", true)
		maybePanic(err)

		// chunk of common fixture, not subject of interest
		author := &git.CommitAttribution{
			Name:  "author",
			Email: "email@domain.wow",
			When:  time.Date(2009, 10, 14, 12, 0, 0, 0, time.UTC),
		}

		Convey("Committing simple content works", func() {
			commitID := createCommit(repo, &git.Commit{
				Author:    author,
				Committer: author,
				Message:   "log message",
				Parents:   nil,
			}, []treeEntry{
				{Filename: "thefile", Filemode: libgit.FilemodeBlob, Content: []byte("hello, world!\n")},
			})
			expectedCommitID := "5409e1f57cf0ffe7a542e78a1c69ae715f2d2abc"
			So(string(commitID), ShouldResemble, expectedCommitID)

			Convey("execgit believe our work", testutil.Requires(
				testutil.RequiresTestLabel("hostgit"),
				func() {
					// check file content retrievable by commit hash
					So(
						execgit.Bake(gosh.Opts{Cwd: "repo"}, "show", expectedCommitID+":thefile").Output(),
						ShouldResemble,
						"hello, world!\n",
					)
				},
			))

			Convey("Creating a new branch at that commit works", func() {
				// create branch
				setBranch(repo, "branchname", commitID)

				Convey("execgit believe our work", testutil.Requires(
					testutil.RequiresTestLabel("hostgit"),
					func() {
						// check branch reference visible
						So(
							execgit.Bake("ls-remote", "repo").Output(),
							ShouldResemble,
							expectedCommitID+"\trefs/heads/branchname\n",
						)
					},
				))
			})
		})
	}))
}

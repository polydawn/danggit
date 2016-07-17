package gat

import (
	"testing"
	"time"

	libgit "github.com/libgit2/git2go"
	. "github.com/smartystreets/goconvey/convey"

	"polydawn.net/danggit/api"
	"polydawn.net/danggit/lib/testutil"
)

func TestHeads(t *testing.T) {
	Convey("Given a local git repo", t, testutil.WithTmpdir(func() {
		repo, err := libgit.InitRepository("repo", true)
		maybePanic(err)

		Convey("which is empty", func() {
			Convey("ListHeads should work", func() {
				resp := ListHeads(git.ReqListHeads{Repo: "repo"})
				So(resp.Error, ShouldBeNil)
				So(resp.Heads, ShouldHaveLength, 0)
			})
		})

		// chunk of common fixture, not subject of interest
		author := &git.CommitAttribution{
			Name:  "author",
			Email: "email@domain.wow",
			When:  time.Date(2009, 10, 14, 12, 0, 0, 0, time.UTC),
		}

		Convey("and given some commits and branches", func() {
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

			// create branch
			setBranch(repo, "branchname", commitID)

			Convey("ListHeads should work", func() {
				resp := ListHeads(git.ReqListHeads{Repo: "repo"})
				So(resp.Error, ShouldBeNil)
				//So(resp.Heads, ShouldHaveLength, 1)
			})
		})
	}))
}

package gat

import (
	"io/ioutil"
	"testing"
	"time"

	libgit "github.com/libgit2/git2go"
	. "github.com/smartystreets/goconvey/convey"

	"polydawn.net/danggit/api"
	"polydawn.net/danggit/lib/testutil"
)

func TestRefs(t *testing.T) {
	Convey("Given a local git repo", t, testutil.WithTmpdir(func() {
		repo, err := libgit.InitRepository("repo", true)
		maybePanic(err)

		Convey("which is empty", func() {
			Convey("ListRefs should work", func() {
				resp := ListRefs(git.ReqListRefs{Repo: "repo"})
				So(resp.Error, ShouldBeNil)
				So(resp.Refs, ShouldHaveLength, 0)
			})

			SkipConvey("ListRefs_Remote should work", func() {
				resp := ListRefs_Remote(git.ReqListRefsRemote{Repo: "repo"})
				So(resp.Error, ShouldBeNil)
				So(resp.Refs, ShouldHaveLength, 0)
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

			Convey("ListRefs should work", func() {
				resp := ListRefs(git.ReqListRefs{Repo: "repo"})
				So(resp.Error, ShouldBeNil)
				So(resp.Refs, ShouldHaveLength, 1)
				So(resp.Refs[0], ShouldResemble, git.Ref{"refs/heads/branchname", "5409e1f57cf0ffe7a542e78a1c69ae715f2d2abc"})
			})

			SkipConvey("ListRefs_Remote should work", func() {
				resp := ListRefs_Remote(git.ReqListRefsRemote{Repo: "repo"})
				So(resp.Error, ShouldBeNil)
				So(resp.Refs, ShouldHaveLength, 1)
			})
		})
	}))
}

func TestRefsIntegration(t *testing.T) {
	Convey("Given a cgit-spawned repo", t, testutil.Requires(
		testutil.RequiresTestLabel("hostgit"),
		testutil.WithTmpdir(func() {
			execgit.Bake("init").RunAndReport()
			ioutil.WriteFile("whop", []byte("woop"), 0644)
			execgit.Bake("add", ".").RunAndReport()
			execgit.Bake("commit", "-mmessage", execgitcommitheaders).RunAndReport()

			Convey("which is just one branch", func() {
				Convey("ListRefs should work", func() {
					resp := ListRefs(git.ReqListRefs{Repo: ".git"})
					So(resp.Error, ShouldBeNil)
					So(resp.Refs, ShouldHaveLength, 1)
					// note that our ListRefs function does *not* return 'HEAD'.
					So(resp.Refs[0], ShouldResemble, git.Ref{"refs/heads/master", "c3d2aa879b52d68570d1b16c448c981e8041c2dd"})
				})
			})
		}),
	))
}

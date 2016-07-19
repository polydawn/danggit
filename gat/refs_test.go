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
				Message:   "log message\n",
				Parents:   nil,
			}, []treeEntry{
				{Filename: "thefile", Filemode: libgit.FilemodeBlob, Content: []byte("hello, world!\n")},
			})
			expectedCommitID := git.CommitID("c57900257a062ce9cd0c69a05bd5837e437626c8")
			So(commitID, ShouldResemble, expectedCommitID)

			// create branch
			setBranch(repo, "branchname", commitID)

			Convey("ListRefs should work", func() {
				resp := ListRefs(git.ReqListRefs{Repo: "repo"})
				So(resp.Error, ShouldBeNil)
				So(resp.Refs, ShouldHaveLength, 1)
				So(resp.Refs[0], ShouldResemble, git.Ref{"refs/heads/branchname", expectedCommitID})
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
			ioutil.WriteFile("thefile", []byte("hello, world!\n"), 0644)
			execgit.Bake("add", ".").RunAndReport()
			execgit.Bake("commit", "-mlog message", execgitcommitheaders).RunAndReport()

			Convey("which is just one branch", func() {
				Convey("ListRefs should work", func() {
					resp := ListRefs(git.ReqListRefs{Repo: ".git"})
					So(resp.Error, ShouldBeNil)
					So(resp.Refs, ShouldHaveLength, 1)
					// note that our ListRefs function does *not* return 'HEAD'.
					So(resp.Refs[0], ShouldResemble, git.Ref{"refs/heads/master", "c57900257a062ce9cd0c69a05bd5837e437626c8"})
				})
			})
		}),
	))
}

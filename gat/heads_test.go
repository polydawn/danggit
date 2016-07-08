package gat

import (
	"testing"

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

		Convey("and given some commits and branches", func() {
			// TODO
			_ = repo

			Convey("ListHeads should work", func() {
				resp := ListHeads(git.ReqListHeads{Repo: "repo"})
				So(resp.Error, ShouldBeNil)
			})
		})
	}))
}

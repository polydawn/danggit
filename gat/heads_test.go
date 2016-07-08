package gat

import (
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"polydawn.net/danggit/lib/testutil"
)

func TestHeads(t *testing.T) {
	Convey("Given a local git repo", t,
		testutil.WithTmpdir(func(c C) {
			mustCmd(execGit("init", "--", "repo-a"))
			var commitHash_1, commitHash_2 string
			testutil.UsingDir("repo-a", func() {
				mustCmd(execGit("commit", "--allow-empty", "-m", "testrepo-a initial commit"))
				commitHash_1 = strings.Trim(mustCmd(execGit("rev-parse", "HEAD")), "\n")
				ioutil.WriteFile("file-a", []byte("abcd"), 0644)
				mustCmd(execGit("add", "."))
				mustCmd(execGit("commit", "-m", "testrepo-a commit 1"))
				commitHash_2 = strings.Trim(mustCmd(execGit("rev-parse", "HEAD")), "\n")
			})

			Convey("Magic should happen", FailureContinues, func() {
			})
		}),
	)
}

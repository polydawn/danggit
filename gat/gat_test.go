package gat

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"polydawn.net/danggit/lib/testutil"
)

func execGit(args ...string) (string, int) {
	cmd := exec.Command("git", args...)
	// ye standard attempts to armour git against bizzarotron system config and global effects
	cmd.Env = []string{
		"GIT_CONFIG_NOSYSTEM=true",
		"HOME=/dev/null",
		"GIT_ASKPASS=/bin/true",
	}
	// slurp all the outputs.  this isn't for purity and grace
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	// have i mentioned how pleasant stdlib exec is at giving exit codes
	exit, err := cmd.Process.Wait()
	if err != nil {
		panic(err)
	}
	waitStatus := exit.Sys().(syscall.WaitStatus)
	if waitStatus.Exited() {
		return string(out), waitStatus.ExitStatus()
	} else {
		panic(waitStatus)
	}
	return string(out), 0
}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func mustCmd(out string, exitCode int) string {
	if exitCode == 0 {
		return out
	}
	panic(fmt.Errorf(
		"expected command to pass; got code %d and message:\n\t%q",
		exitCode, out,
	))
}

func TestHello(t *testing.T) {
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

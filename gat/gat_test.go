package gat

import (
	"fmt"
	"os/exec"
	"syscall"
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

func mustCmd(out string, exitCode int) string {
	if exitCode == 0 {
		return out
	}
	panic(fmt.Errorf(
		"expected command to pass; got code %d and message:\n\t%q",
		exitCode, out,
	))
}

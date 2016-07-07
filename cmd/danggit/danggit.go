package main

import (
	"fmt"
	"os"

	"github.com/libgit2/git2go"
)

func main() {
	repo, err := git.OpenRepository(".")
	if err != nil {
		fmt.Printf("not a repo: %s\n", err)
		os.Exit(5)
	}

	desc, err := repo.DescribeWorkdir(&git.DescribeOptions{ShowCommitOidAsFallback: true})
	if err != nil {
		fmt.Printf("an ineffable miracle beyond description: %s\n", err)
		os.Exit(6)
	}
	fmt.Printf("repo: %s\n", desc)
}

/*
	Hokay, so.

	Is it implementation detail or showing truth to have the api take a batch of operations on a repo in one bulk message.

	- If the `OpenRepository` call makes a lock on the index, it is truth.
	- If it doesn't, then batching things should very well be transparent.
	- If the error paths may change based on whether or not your request required a fresh open, that would also indicate explicit batching is a necessary truth.

	Observations:

	- still haven't detected boolean answer to that last
	- there are definitely operations that don't operate on any repo at all, so keep that in mind
	- mind the iterables, and the progress bars.  this will definitely not be an rpc api that's 1:1 on ask/answer messages.
*/

/*
	Some attitudes we have:

	Repo data dirs and checked out working trees are *always* treated separately.
	We will not autodetect on these things.  If you want the git data dir at '$WORKTREE/.git/', *say so*.
	We believe this it the right thing to do because many, many git operations desired in scripting we've performed
	become radically simpler, more predictable, and less error-prone if you simply do them in a clean working
	tree (while also understanding that you *can do this* without giving up the shared cache of git data).
	Pivoting your understanding of git to regard the data dir as the center of the universe and worktrees
	as an incidental, optional, no-blessed-default thing will improve your designs significantly.

	It should always be extremely clear whether you're talking about a local filesystem repo or a remote repo.
	If your application doesn't know whether it's talking about remote resources or local resources, it's poorly
	designed: it's going to suffer from the non-transparent performance implications of that; and any local-only
	code paths will suffer from complex error handling for cases that can't actually happen.
	Therefore, our API always emphasizes functions that operate on local-only paths; functions that may accept
	remote repo paths are always the exception, and have different names to clear mark them as such.
	Functions that operation on remote paths may also accept local paths, but never vice versa.
*/

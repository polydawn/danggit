package gat

import (
	libgit "github.com/libgit2/git2go"

	"polydawn.net/danggit/api"
)

func ListHeads(pth git.LocalRepoPath) []git.Head {
	return nil // TODO
}

func ListHeads_Remote(req git.ReqListHeadsRemote) git.RespListHeads {
	remote := &libgit.Remote{}
	// dial
	err := remote.ConnectFetch(nil, nil, nil)
	if err != nil {
		return git.RespListHeads{Error: err}
	}
	// read heads
	theirHeads, err := remote.Ls(req.Filters...)
	if err != nil {
		return git.RespListHeads{Error: err}
	}
	// flip to our api types
	heads := make([]git.Head, len(theirHeads))
	for i, head := range theirHeads {
		heads[i] = git.Head{
			RefName:  head.Name,
			CommitID: head.Id.String(),
		}
	}
	return git.RespListHeads{Heads: heads}
}

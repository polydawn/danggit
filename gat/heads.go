package gat

import (
	libgit "github.com/libgit2/git2go"

	"polydawn.net/danggit/api"
)

func ListHeads(req git.ReqListHeads) git.RespListHeads {
	// dial
	repo, err := openRepository(req.Repo)
	if err != nil {
		return git.RespListHeads{Error: err}
	}
	// read heads
	// FIXME why the hell does this take one glob when `remote.Ls` takes a filter *list*...
	itr, err := repo.NewReferenceIteratorGlob("sightodo")
	if err != nil {
		return git.RespListHeads{Error: err}
	}
	// flip to our api types
	heads := make([]git.Head, 0)
	i := 0
	for ref, err := itr.Next(); err == nil; ref, err = itr.Next() {
		// TODO maybe check 'isBranch' and '!isRemote' here?  need to decide what semantics we really intend here
		// FIXME ... which makes me realize, there's a reason heads and refs are different words; these funcs might not be using them entirely wisely
		heads[i] = git.Head{
			RefName:  ref.Name(),
			CommitID: ref.Target().String(),
		}
		i++
	}
	return git.RespListHeads{Heads: heads}
}

func ListHeads_Remote(req git.ReqListHeadsRemote) git.RespListHeads {
	remote := &libgit.Remote{}
	// dial
	err := remote.ConnectFetch(nil, nil, nil) // FIXME wait, what?  where do i set the path?
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

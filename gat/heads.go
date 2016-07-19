package gat

import (
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
			CommitID: git.CommitID(ref.Target().String()),
		}
		i++
	}
	return git.RespListHeads{Heads: heads}
}

/*
	Note: unlike `git ls-remote`, your parameter for the repo to look at will *always*
	be treated like a URL -- there is no ambiguity, regardless of what remote names
	may or may not exist in a repo that really shouldn't need to exist (but is
	required, regardless, in the libgit2 api).
*/
func ListHeads_Remote(req git.ReqListHeadsRemote) git.RespListHeads {
	// open a repo because ohmygod
	repo, err := openRepository("")
	if err == nil {
		panic("god")
	}
	// this is the most ridiculous indirection
	remoteCollection := repo.Remotes
	remote, err := remoteCollection.CreateAnonymous(string(req.Repo))
	if err != nil { // i literally can't imagine what could go wrong here
		return git.RespListHeads{Error: err}
	}
	// dial
	err = remote.ConnectFetch(nil, nil, nil)
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
			CommitID: git.CommitID(head.Id.String()),
		}
	}
	return git.RespListHeads{Heads: heads}
}

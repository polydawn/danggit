package gat

import (
	"polydawn.net/danggit/api"
)

func ListRefs(req git.ReqListRefs) git.RespListRefs {
	// dial
	repo, err := openRepository(req.Repo)
	if err != nil {
		return git.RespListRefs{Error: err}
	}
	// read heads
	// FIXME why the hell does this take one glob when `remote.Ls` takes a filter *list*...
	itr, err := repo.NewReferenceIteratorGlob("*")
	if err != nil {
		return git.RespListRefs{Error: err}
	}
	// flip to our api types
	refs := make([]git.Ref, 0)
	for cref, err := itr.Next(); err == nil; cref, err = itr.Next() {
		// TODO maybe check 'isBranch' and '!isRemote' here?  need to decide what semantics we really intend here
		refs = append(refs, git.Ref{
			Name:  cref.Name(),
			CommitID: git.CommitID(cref.Target().String()),
		})
	}
	return git.RespListRefs{Refs: refs}
}

/*
	Note: unlike `git ls-remote`, your parameter for the repo to look at will *always*
	be treated like a URL -- there is no ambiguity, regardless of what remote names
	may or may not exist in a repo that really shouldn't need to exist (but is
	required, regardless, in the libgit2 api).
*/
func ListRefs_Remote(req git.ReqListRefsRemote) git.RespListRefs {
	// open a repo because ohmygod
	repo, err := openRepository("")
	if err == nil {
		panic("god")
	}
	// this is the most ridiculous indirection
	remoteCollection := repo.Remotes
	remote, err := remoteCollection.CreateAnonymous(string(req.Repo))
	if err != nil { // i literally can't imagine what could go wrong here
		return git.RespListRefs{Error: err}
	}
	// dial
	err = remote.ConnectFetch(nil, nil, nil)
	if err != nil {
		return git.RespListRefs{Error: err}
	}
	// read heads
	theirRefs, err := remote.Ls(req.Filters...)
	if err != nil {
		return git.RespListRefs{Error: err}
	}
	// flip to our api types
	refs := make([]git.Ref, len(theirRefs))
	for i, cref := range theirRefs {
		refs[i] = git.Ref{
			Name:  cref.Name,
			CommitID: git.CommitID(cref.Id.String()),
		}
	}
	return git.RespListRefs{Refs: refs}
}

package git

import (
	"time"
)

/*
	Note that many structures in this package are very close analogs of things
	also in https://godoc.org/github.com/libgit2/git2go -- for example,
	our `CommitAttribution` type is almost identical to `git2go.Signature`.
	This is convergent evolution: we're expressing the same things; but it's
	important for us to copy them here, both so that the symbols are easy to
	link in other go programs while staying a lightyear away from cgo,
	and also because it makes it a helluva lot easier for us to write our
	serialization layer.
*/

type LocalRepoPath string
type RemoteRepoPath string

type Call string

const (
	Call_ListHeads       = Call("ListHeads")       // ReqListHeads -> RespListHeads
	Call_ListHeadsRemote = Call("ListHeadsRemote") // ReqListHeadsRemote -> RespListHeads
)

// REVIEW: we can't export libgit.Oid up here, but perhaps we should at least maintain their byte format
//  Because the hex parsing and error checking and etc that libgit.NewOid performs is ridiculous.
//  `libgit.NewOidFromBytes` does a copy but that feels like an acceptable rounding error for our API.

type CommitID string
type TreeID string

type Req struct {
	ThreadID string
	Call     Call
	Params   interface{}
}

type Resp struct {
	ThreadID string
	Params   interface{}
}

type Head struct {
	RefName  string
	CommitID CommitID
}

type Commit struct {
	Author    *CommitAttribution
	Committer *CommitAttribution
	Message   string
	Parents   []CommitID
	TreeID    TreeID
}

type CommitAttribution struct {
	Name  string
	Email string
	When  time.Time
}

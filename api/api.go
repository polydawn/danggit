package git

type LocalRepoPath string
type RemoteRepoPath string

type Call string

const (
	Call_ListHeads       = Call("ListHeads")       // ReqListHeads -> RespListHeads
	Call_ListHeadsRemote = Call("ListHeadsRemote") // ReqListHeadsRemote -> RespListHeads
)

type Req struct {
	ThreadID string
	Call     Call
	Params   interface{}
}

type Resp struct {
	ThreadID string
	Params   interface{}
}

type ReqListHeads struct {
	Repo    LocalRepoPath
	Filters []string
}

type ReqListHeadsRemote struct {
	Repo    LocalRepoPath
	Filters []string
}

type RespListHeads struct {
	Heads []Head
	Error error
}

type Head struct {
	RefName  string
	CommitID string
}

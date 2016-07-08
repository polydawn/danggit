package git

type LocalRepoPath string

type ReqListHeads struct {
	ThreadID string
	Repo     LocalRepoPath
	Filter   string
}

type RespListHeads struct {
	ThreadID string
	Heads    []Head
	Error    error
}

type Head struct {
	RefName  string
	CommitID string
}

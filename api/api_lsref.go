package git

type ReqListHeads struct {
	Repo    LocalRepoPath
	Filters []string
}

type ReqListHeadsRemote struct {
	Repo    RemoteRepoPath
	Filters []string
}

type RespListHeads struct {
	Heads []Head
	Error error
}

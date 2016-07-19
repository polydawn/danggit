package git

type ReqListRefs struct {
	Repo    LocalRepoPath
	Filters []string
}

type ReqListRefsRemote struct {
	Repo    RemoteRepoPath
	Filters []string
}

type RespListRefs struct {
	Refs []Ref
	Error error
}

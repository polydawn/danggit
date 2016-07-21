package git

type (
	DoListRefs       func(req ReqListRefs) RespListRefs
	DoListRefsRemote func(req ReqListRefsRemote) RespListRefs
)

type ReqListRefs struct {
	Repo    LocalRepoPath
	Filters []string
}

type ReqListRefsRemote struct {
	Repo    RemoteRepoPath
	Filters []string
}

type RespListRefs struct {
	Refs  []Ref
	Error error
}

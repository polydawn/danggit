package git

type Call string

const (
	Call_ListRefs       = Call("ListRefs")       // ReqListRefs -> RespListRefs
	Call_ListRefsRemote = Call("ListRefsRemote") // ReqListRefsRemote -> RespListRefs
)

var table = []struct {
	Call
	ReqType interface{}
	Route   func(d *Dispatcher) interface{}
}{
	{Call_ListRefs, &ReqListRefs{}, func(d *Dispatcher) interface{} { return d.DoListRefs }},
}

func (req *Req) InitCallParams() {
	req.Params = func() interface{} {
		switch req.Call {
		case Call_ListRefs:
			return &ReqListRefs{}
		case Call_ListRefsRemote:
			return &ReqListRefsRemote{}
		default:
			return nil
		}
	}()
}

type Dispatcher struct {
	DoListRefs       DoListRefs
	DoListRefsRemote DoListRefsRemote
}

func (d *Dispatcher) DispatchReq(req *Req) {
	switch req.Call {
	case Call_ListRefs:
		d.DoListRefs(*(req.Params.(*ReqListRefs))) // TODO serialize and send return
	case Call_ListRefsRemote:
		d.DoListRefsRemote(*(req.Params.(*ReqListRefsRemote))) // TODO serialize and send return
	default:
		panic("unknown call")
	}
}

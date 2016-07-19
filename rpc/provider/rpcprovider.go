package rpcprov

import (
	"encoding/json"
	"io"

	"polydawn.net/danggit/api"
)

func Main(stdin io.Reader, sink chan<- *git.Req) {
	defer close(sink)
	dec := json.NewDecoder(stdin)
	for {
		req := &git.Req{
			Params: &json.RawMessage{},
		}
		switch err := dec.Decode(&req); err {
		case io.EOF:
			return
		case nil:
			// pass
		default:
			panic(err)
		}
		raw := []byte(*req.Params.(*json.RawMessage))
		req.InitCallParams()
		if err := json.Unmarshal(raw, req.Params); err != nil {
			panic(err)
		}
		sink <- req
	}
}

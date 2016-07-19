package rpcprov

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"polydawn.net/danggit/api"
)

func TestReqReader(t *testing.T) {
	Convey("Given a wow", t, func() {
		input := bytes.NewBuffer([]byte(`{"ThreadID":"wow", "Call":"ListRefs", "Params":{"Repo":"."}}`))
		output := make(chan *git.Req, 2)
		Main(input, output)
		So((<-output).Params, ShouldResemble, &git.ReqListRefs{Repo: "."})
	})
}

package agent

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
)

func (p *WebServer) ctrMain(ir *iron.Request) {
	ir.ApiOutput("hello world", snettypes.CODE_OK, "")
}

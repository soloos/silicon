package soloboatsvr

import (
	"soloos/common/iron"
)

func (p *WebServer) ctrMain(ir *iron.Request) {
	ir.Redirect("/SDFS/NameNode")
}

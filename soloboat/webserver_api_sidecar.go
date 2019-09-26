package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/soloboat/sidecartypes"
)

func (p *WebServer) apiSidecarHeartBeat(ir *iron.Request) {
	var (
		heartbeat sidecartypes.SidecarHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	p.soloboatIns.sideCarDriver.SidecarHeartBeat(heartbeat)
	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

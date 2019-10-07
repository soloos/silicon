package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snet"
	"soloos/soloboat/sidecartypes"
)

func (p *WebServer) apiSidecarHeartBeat(ir *iron.Request) {
	var (
		heartbeat sidecartypes.SidecarHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snet.CODE_502, err.Error())
		return
	}

	p.soloboatIns.sideCarDriver.SidecarHeartBeat(heartbeat)
	ir.ApiOutput(nil, snet.CODE_OK, "heartbeat success")
}

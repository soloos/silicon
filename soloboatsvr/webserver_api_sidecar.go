package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/soloboat/sidecartypes"
)

func (p *WebServer) apiSideCarHeartBeat(ir *iron.Request) {
	var (
		heartbeat sidecartypes.SideCarHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	p.soloBoatSvr.sideCarDriver.SideCarHeartBeat(heartbeat)
	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

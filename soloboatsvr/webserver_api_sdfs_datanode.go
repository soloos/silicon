package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
)

func (p *WebServer) apiSDFSDataNodeHeartBeat(ir *iron.Request) {
	var (
		heartbeat sdfsapitypes.DataNodeHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	p.soloBoatSvr.sdfsDataNodeDriver.SDFSDataNodeHeartBeat(heartbeat)
	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

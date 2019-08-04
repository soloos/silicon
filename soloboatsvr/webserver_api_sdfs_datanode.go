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

	err = p.soloBoatSvr.sdfsDataNodeDriver.SDFSDataNodeHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

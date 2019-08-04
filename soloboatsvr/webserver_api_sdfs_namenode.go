package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
)

func (p *WebServer) apiSDFSNameNodeHeartBeat(ir *iron.Request) {
	var (
		heartbeat sdfsapitypes.NameNodeHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	err = p.soloBoatSvr.sdfsNameNodeDriver.SDFSNameNodeHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/sdbapitypes"
	"soloos/common/snettypes"
)

func (p *WebServer) apiSDBOneHeartBeat(ir *iron.Request) {
	var (
		heartbeat sdbapitypes.SDBOneHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	err = p.soloBoatSvr.sdbOneDriver.SDBOneHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

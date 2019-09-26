package soloboat

import (
	"soloos/common/iron"
	"soloos/common/solodbapitypes"
	"soloos/common/snettypes"
)

func (p *WebServer) apiSolodbHeartBeat(ir *iron.Request) {
	var (
		heartbeat solodbapitypes.SolodbHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	err = p.soloboatIns.solodbDriver.SolodbHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

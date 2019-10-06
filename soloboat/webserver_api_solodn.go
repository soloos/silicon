package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/common/solofsapitypes"
)

func (p *WebServer) apiSolodnHeartBeat(ir *iron.Request) {
	var (
		heartbeat solofsapitypes.SolodnHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	err = p.soloboatIns.solodnDriver.SolodnHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

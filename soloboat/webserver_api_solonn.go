package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/common/solofsapitypes"
)

func (p *WebServer) apiSolonnHeartBeat(ir *iron.Request) {
	var (
		heartbeat solofsapitypes.SolonnHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	err = p.soloboatIns.solonnDriver.SolonnHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

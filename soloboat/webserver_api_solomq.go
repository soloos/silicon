package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/common/solomqapitypes"
)

func (p *WebServer) apiSolomqHeartBeat(ir *iron.Request) {
	var (
		heartbeat solomqapitypes.SolomqHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	err = p.soloboatIns.solomqSolomqDriver.SolomqHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

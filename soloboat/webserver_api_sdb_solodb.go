package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snet"
	"soloos/common/solodbapitypes"
)

func (p *WebServer) apiSolodbHeartBeat(ir *iron.Request) {
	var (
		heartbeat solodbapitypes.SolodbHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snet.CODE_502, err.Error())
		return
	}

	err = p.soloboatIns.solodbDriver.SolodbHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snet.CODE_502, err.Error())
		return
	}

	ir.ApiOutput(nil, snet.CODE_OK, "heartbeat success")
}

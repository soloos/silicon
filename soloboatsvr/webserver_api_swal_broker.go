package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/common/swalapitypes"
)

func (p *WebServer) apiSWALBrokerHeartBeat(ir *iron.Request) {
	var (
		heartbeat swalapitypes.BrokerHeartBeat
		err       error
	)

	err = ir.DecodeBodyJSONData(&heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	err = p.soloBoatSvr.swalBrokerDriver.SWALBrokerHeartBeat(heartbeat)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "heartbeat success")
}

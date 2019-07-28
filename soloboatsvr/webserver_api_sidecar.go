package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/log"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboatsdk"
	"soloos/soloboat/soloboattypes"
)

func (p *WebServer) apiSideCarHeartBeat(ir *iron.Request) {
	var (
		req soloboatsdk.SideCarInfoReqJSON
		err error
	)

	err = ir.DecodeBodyJSONData(&req)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	var sideCarInfo = soloboattypes.DecodeSideCarInfoJSON(req.SideCarInfoJSON)
	err = p.soloBoatSvr.sideCarDriver.SidCarHeartBeat(sideCarInfo)
	log.Error("fuck register", req.PeerID, err)
	if err != nil {
		return
	}

	ir.ApiOutput(nil, snettypes.CODE_OK, "")
}

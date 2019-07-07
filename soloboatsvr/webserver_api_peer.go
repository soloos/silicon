package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/log"
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"soloos/soloboat/soloboatsdk"
)

func (p *WebServer) apiPeerList(ir *iron.Request) {
	var ret []snettypes.PeerJSON
	p.soloBoatSvr.SNetDriver.ListPeer(func(uObj offheap.LKVTableObjectUPtrWithBytes64) bool {
		var peer = *snettypes.PeerUintptr(uObj).Ptr()
		ret = append(ret, snettypes.PeerToPeerJSON(peer))
		return true
	})
	SortSNetPeerJSON(ret)
	ir.ApiOutput(ret, snettypes.CODE_OK, "")
}

func (p *WebServer) apiPeerHeartBeat(ir *iron.Request) {
	var (
		req soloboatsdk.HeartBeatReqJSON
		err error
	)

	err = ir.DecodeBodyJSONData(&req)
	if err != nil {
		ir.ApiOutput(nil, snettypes.CODE_502, err.Error())
		return
	}

	log.Error("fuck heatbeat", req.PeerID)
	ir.ApiOutput(nil, snettypes.CODE_OK, "")
}

package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
)

func (p *WebServer) prepareCtrSNetPeer(ir *iron.Request) bool {
	var module = ViewModule{
		Name: "SNetPeer",
		CH:   "SNetPeer",
		URL:  "/SNet/Peer",
	}
	ir.ViewData["Module"] = module
	return true
}

func (p *WebServer) ctrSNetPeer(ir *iron.Request) {
	var ret []snettypes.PeerJSON
	p.soloBoatSvr.SNetDriver.ListPeer(func(uObj offheap.LKVTableObjectUPtrWithBytes64) bool {
		var peer = *snettypes.PeerUintptr(uObj).Ptr()
		ret = append(ret, snettypes.PeerToPeerJSON(peer))
		return true
	})
	ir.ViewData["SNetPeerList"] = ret
	ir.Render("/SNet/Peer/Index")
}

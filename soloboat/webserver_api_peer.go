package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/solodb/offheap"
)

func (p *WebServer) apiPeerList(ir *iron.Request) {
	var ret []snettypes.PeerJSON
	p.soloboatIns.SNetDriver.ListPeer(func(uObj offheap.LKVTableObjectUPtrWithBytes64) bool {
		var uptr = snettypes.PeerUintptr(uObj)
		var peer = *uptr.Ptr()
		ret = append(ret, snettypes.PeerToPeerJSON(peer))
		return true
	})
	SortSNetPeerJSON(ret)
	ir.ApiOutput(ret, snettypes.CODE_OK, "")
}

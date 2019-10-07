package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snet"
	"soloos/solodb/offheap"
)

func (p *WebServer) apiPeerList(ir *iron.Request) {
	var ret []snet.PeerJSON
	p.soloboatIns.SNetDriver.ListPeer(func(uObj offheap.LKVTableObjectUPtrWithBytes64) bool {
		var uptr = snet.PeerUintptr(uObj)
		var peer = *uptr.Ptr()
		ret = append(ret, snet.PeerToPeerJSON(peer))
		return true
	})
	SortSNetPeerJSON(ret)
	ir.ApiOutput(ret, snet.CODE_OK, "")
}

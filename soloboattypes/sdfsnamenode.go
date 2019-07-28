package soloboattypes

import (
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"unsafe"
)

const (
	SDFSNameNodeInfoStructSize = unsafe.Sizeof(SDFSNameNodeInfo{})
)

type SDFSNameNodeInfoJSON struct {
	PeerID string
}

func DecodeSDFSNameNodeInfoJSON(sideCarInfoJSON SDFSNameNodeInfoJSON) SDFSNameNodeInfo {
	var ret SDFSNameNodeInfo
	copy(ret.ID[:], []byte(sideCarInfoJSON.PeerID))
	return ret
}

type SDFSNameNodeInfoUintptr uintptr

func (u SDFSNameNodeInfoUintptr) Ptr() *SDFSNameNodeInfo {
	return (*SDFSNameNodeInfo)(unsafe.Pointer(u))
}

type SDFSNameNodeInfo struct {
	offheap.LKVTableObjectWithBytes64 `db:"-"`
}

func (p *SDFSNameNodeInfo) PeerID() snettypes.PeerID { return snettypes.PeerID(p.ID) }

func (p *SDFSNameNodeInfo) PeerIDStr() string { return snettypes.PeerID(p.ID).Str() }

package soloboattypes

import (
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"unsafe"
)

const (
	SideCarInfoStructSize = unsafe.Sizeof(SideCarInfo{})
)

type SideCarInfoJSON struct {
	PeerID string
}

func DecodeSideCarInfoJSON(sideCarInfoJSON SideCarInfoJSON) SideCarInfo {
	var ret SideCarInfo
	copy(ret.ID[:], []byte(sideCarInfoJSON.PeerID))
	return ret
}

type SideCarInfoUintptr uintptr

func (u SideCarInfoUintptr) Ptr() *SideCarInfo { return (*SideCarInfo)(unsafe.Pointer(u)) }

type SideCarInfo struct {
	offheap.LKVTableObjectWithBytes64 `db:"-"`
}

func (p *SideCarInfo) PeerID() snettypes.PeerID { return snettypes.PeerID(p.ID) }

func (p *SideCarInfo) PeerIDStr() string { return snettypes.PeerID(p.ID).Str() }

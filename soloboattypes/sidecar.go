package soloboattypes

import (
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"soloos/soloboat/sidecartypes"
	"time"
	"unsafe"
)

const (
	SideCarInfoStructSize = unsafe.Sizeof(SideCarInfo{})
)

type SideCarInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSideCarInfoJSON(sideCarInfoJSON SideCarInfoJSON) SideCarInfo {
	var ret SideCarInfo
	copy(ret.ID[:], []byte(sideCarInfoJSON.PeerID))
	ret.LastHeatBeatAt = time.Unix(sideCarInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSideCarInfoJSON(sideCarInfo SideCarInfo) SideCarInfoJSON {
	var ret SideCarInfoJSON
	ret.PeerID = string(sideCarInfo.ID[:])
	ret.LastHeatBeatAt = sideCarInfo.LastHeatBeatAt.Unix()
	return ret
}

type SideCarInfoUintptr uintptr

func (u SideCarInfoUintptr) Ptr() *SideCarInfo {
	return (*SideCarInfo)(unsafe.Pointer(u))
}

type SideCarInfo struct {
	offheap.LKVTableObjectWithBytes64 `db:"-"`
	LastHeatBeatAt                    time.Time
	LastHeatBeatAtStr                 string
	SRPCServerAddr                    string
	WebServerAddr                     string
	sidecartypes.SideCarHeartBeat
}

func (p *SideCarInfo) PeerID() snettypes.PeerID { return snettypes.PeerID(p.ID) }

func (p *SideCarInfo) PeerIDStr() string { return snettypes.PeerID(p.ID).Str() }

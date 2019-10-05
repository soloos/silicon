package soloboattypes

import (
	"soloos/common/snettypes"
	"soloos/solodb/offheap"
	"soloos/soloboat/sidecartypes"
	"time"
	"unsafe"
)

const (
	SidecarInfoStructSize = unsafe.Sizeof(SidecarInfo{})
)

type SidecarInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSidecarInfoJSON(sideCarInfoJSON SidecarInfoJSON) SidecarInfo {
	var ret SidecarInfo
	copy(ret.ID[:], []byte(sideCarInfoJSON.PeerID))
	ret.LastHeatBeatAt = time.Unix(sideCarInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSidecarInfoJSON(sideCarInfo SidecarInfo) SidecarInfoJSON {
	var ret SidecarInfoJSON
	ret.PeerID = string(sideCarInfo.ID[:])
	ret.LastHeatBeatAt = sideCarInfo.LastHeatBeatAt.Unix()
	return ret
}

type SidecarInfoUintptr uintptr

func (u SidecarInfoUintptr) Ptr() *SidecarInfo {
	return (*SidecarInfo)(unsafe.Pointer(u))
}

type SidecarInfo struct {
	offheap.LKVTableObjectWithBytes64 `db:"-"`
	LastHeatBeatAt                    time.Time
	LastHeatBeatAtStr                 string
	SrpcServerAddr                    string
	WebServerAddr                     string
	sidecartypes.SidecarHeartBeat
}

func (p *SidecarInfo) PeerID() snettypes.PeerID { return snettypes.PeerID(p.ID) }

func (p *SidecarInfo) PeerIDStr() string { return snettypes.PeerID(p.ID).Str() }

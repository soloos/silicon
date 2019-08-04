package soloboattypes

import (
	"soloos/common/sdbapitypes"
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"time"
	"unsafe"
)

const (
	SDBOneInfoStructSize = unsafe.Sizeof(SDBOneInfo{})
)

type SDBOneInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSDBOneInfoJSON(sdbOneInfoJSON SDBOneInfoJSON) SDBOneInfo {
	var ret SDBOneInfo
	copy(ret.ID[:], []byte(sdbOneInfoJSON.PeerID))
	ret.LastHeatBeatAt = time.Unix(sdbOneInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSDBOneInfoJSON(sdbOneInfo SDBOneInfo) SDBOneInfoJSON {
	var ret SDBOneInfoJSON
	ret.PeerID = string(sdbOneInfo.ID[:])
	ret.LastHeatBeatAt = sdbOneInfo.LastHeatBeatAt.Unix()
	return ret
}

type SDBOneInfoUintptr uintptr

func (u SDBOneInfoUintptr) Ptr() *SDBOneInfo {
	return (*SDBOneInfo)(unsafe.Pointer(u))
}

type SDBOneInfo struct {
	offheap.LKVTableObjectWithBytes64 `db:"-"`
	LastHeatBeatAt                    time.Time
	LastHeatBeatAtStr                 string
	SRPCServerAddr                    string
	WebServerAddr                     string
	sdbapitypes.SDBOneHeartBeat
}

func (p *SDBOneInfo) PeerID() snettypes.PeerID { return snettypes.PeerID(p.ID) }

func (p *SDBOneInfo) PeerIDStr() string { return snettypes.PeerID(p.ID).Str() }

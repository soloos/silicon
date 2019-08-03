package soloboattypes

import (
	"soloos/common/snettypes"
	"soloos/common/swalapitypes"
	"soloos/sdbone/offheap"
	"time"
	"unsafe"
)

const (
	SWALBrokerInfoStructSize = unsafe.Sizeof(SWALBrokerInfo{})
)

type SWALBrokerInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSWALBrokerInfoJSON(swalBrokerInfoJSON SWALBrokerInfoJSON) SWALBrokerInfo {
	var ret SWALBrokerInfo
	copy(ret.ID[:], []byte(swalBrokerInfoJSON.PeerID))
	ret.LastHeatBeatAt = time.Unix(swalBrokerInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSWALBrokerInfoJSON(swalBrokerInfo SWALBrokerInfo) SWALBrokerInfoJSON {
	var ret SWALBrokerInfoJSON
	ret.PeerID = string(swalBrokerInfo.ID[:])
	ret.LastHeatBeatAt = swalBrokerInfo.LastHeatBeatAt.Unix()
	return ret
}

type SWALBrokerInfoUintptr uintptr

func (u SWALBrokerInfoUintptr) Ptr() *SWALBrokerInfo {
	return (*SWALBrokerInfo)(unsafe.Pointer(u))
}

type SWALBrokerInfo struct {
	offheap.LKVTableObjectWithBytes64 `db:"-"`
	LastHeatBeatAt                    time.Time
	LastHeatBeatAtStr                 string
	SRPCServerAddr                    string
	WebServerAddr                     string
	swalapitypes.BrokerHeartBeat
}

func (p *SWALBrokerInfo) PeerID() snettypes.PeerID { return snettypes.PeerID(p.ID) }

func (p *SWALBrokerInfo) PeerIDStr() string { return snettypes.PeerID(p.ID).Str() }

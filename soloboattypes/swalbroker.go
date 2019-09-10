package soloboattypes

import (
	"soloos/common/snettypes"
	"soloos/common/swalapitypes"
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
	ret.PeerID.SetStr(swalBrokerInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(swalBrokerInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSWALBrokerInfoJSON(swalBrokerInfo SWALBrokerInfo) SWALBrokerInfoJSON {
	var ret SWALBrokerInfoJSON
	ret.PeerID = string(swalBrokerInfo.PeerID.Str())
	ret.LastHeatBeatAt = swalBrokerInfo.LastHeatBeatAt.Unix()
	return ret
}

type SWALBrokerInfoUintptr uintptr

func (u SWALBrokerInfoUintptr) Ptr() *SWALBrokerInfo {
	return (*SWALBrokerInfo)(unsafe.Pointer(u))
}

type SWALBrokerInfo struct {
	snettypes.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SRPCServerAddr    string
	WebServerAddr     string
	swalapitypes.BrokerHeartBeat
}

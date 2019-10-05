package soloboattypes

import (
	"soloos/common/snettypes"
	"soloos/common/solomqapitypes"
	"time"
	"unsafe"
)

const (
	SolomqInfoStructSize = unsafe.Sizeof(SolomqInfo{})
)

type SolomqInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSolomqInfoJSON(solomqSolomqInfoJSON SolomqInfoJSON) SolomqInfo {
	var ret SolomqInfo
	ret.PeerID.SetStr(solomqSolomqInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(solomqSolomqInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSolomqInfoJSON(solomqSolomqInfo SolomqInfo) SolomqInfoJSON {
	var ret SolomqInfoJSON
	ret.PeerID = string(solomqSolomqInfo.PeerID.Str())
	ret.LastHeatBeatAt = solomqSolomqInfo.LastHeatBeatAt.Unix()
	return ret
}

type SolomqInfoUintptr uintptr

func (u SolomqInfoUintptr) Ptr() *SolomqInfo {
	return (*SolomqInfo)(unsafe.Pointer(u))
}

type SolomqInfo struct {
	snettypes.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SrpcServerAddr    string
	WebServerAddr     string
	solomqapitypes.SolomqHeartBeat
}

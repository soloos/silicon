package soloboattypes

import (
	"soloos/common/sdbapitypes"
	"soloos/common/snettypes"
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
	ret.PeerID.SetStr(sdbOneInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(sdbOneInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSDBOneInfoJSON(sdbOneInfo SDBOneInfo) SDBOneInfoJSON {
	var ret SDBOneInfoJSON
	ret.PeerID = string(sdbOneInfo.PeerID.Str())
	ret.LastHeatBeatAt = sdbOneInfo.LastHeatBeatAt.Unix()
	return ret
}

type SDBOneInfo struct {
	snettypes.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SRPCServerAddr    string
	WebServerAddr     string
	sdbapitypes.SDBOneHeartBeat
}

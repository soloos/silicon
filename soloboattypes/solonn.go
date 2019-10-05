package soloboattypes

import (
	"soloos/common/solofsapitypes"
	"soloos/common/snettypes"
	"time"
	"unsafe"
)

const (
	SolonnInfoStructSize = unsafe.Sizeof(SolonnInfo{})
)

type SolonnInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSolonnInfoJSON(solonnInfoJSON SolonnInfoJSON) SolonnInfo {
	var ret SolonnInfo
	ret.PeerID.SetStr(solonnInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(solonnInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSolonnInfoJSON(solonnInfo SolonnInfo) SolonnInfoJSON {
	var ret SolonnInfoJSON
	ret.PeerID = string(solonnInfo.PeerID.Str())
	ret.LastHeatBeatAt = solonnInfo.LastHeatBeatAt.Unix()
	return ret
}

type SolonnInfo struct {
	snettypes.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SrpcServerAddr    string
	WebServerAddr     string
	solofsapitypes.SolonnHeartBeat
}

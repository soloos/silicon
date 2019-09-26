package soloboattypes

import (
	"soloos/common/solodbapitypes"
	"soloos/common/snettypes"
	"time"
	"unsafe"
)

const (
	SolodbInfoStructSize = unsafe.Sizeof(SolodbInfo{})
)

type SolodbInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSolodbInfoJSON(solodbInfoJSON SolodbInfoJSON) SolodbInfo {
	var ret SolodbInfo
	ret.PeerID.SetStr(solodbInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(solodbInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSolodbInfoJSON(solodbInfo SolodbInfo) SolodbInfoJSON {
	var ret SolodbInfoJSON
	ret.PeerID = string(solodbInfo.PeerID.Str())
	ret.LastHeatBeatAt = solodbInfo.LastHeatBeatAt.Unix()
	return ret
}

type SolodbInfo struct {
	snettypes.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SRPCServerAddr    string
	WebServerAddr     string
	solodbapitypes.SolodbHeartBeat
}

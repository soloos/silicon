package soloboattypes

import (
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
	"time"
	"unsafe"
)

const (
	SDFSNameNodeInfoStructSize = unsafe.Sizeof(SDFSNameNodeInfo{})
)

type SDFSNameNodeInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSDFSNameNodeInfoJSON(sdfsNameNodeInfoJSON SDFSNameNodeInfoJSON) SDFSNameNodeInfo {
	var ret SDFSNameNodeInfo
	ret.PeerID.SetStr(sdfsNameNodeInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(sdfsNameNodeInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSDFSNameNodeInfoJSON(sdfsNameNodeInfo SDFSNameNodeInfo) SDFSNameNodeInfoJSON {
	var ret SDFSNameNodeInfoJSON
	ret.PeerID = string(sdfsNameNodeInfo.PeerID.Str())
	ret.LastHeatBeatAt = sdfsNameNodeInfo.LastHeatBeatAt.Unix()
	return ret
}

type SDFSNameNodeInfo struct {
	snettypes.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SRPCServerAddr    string
	WebServerAddr     string
	sdfsapitypes.NameNodeHeartBeat
}

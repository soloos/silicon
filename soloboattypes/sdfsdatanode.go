package soloboattypes

import (
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
	"time"
	"unsafe"
)

const (
	SDFSDataNodeInfoStructSize = unsafe.Sizeof(SDFSDataNodeInfo{})
)

type SDFSDataNodeInfoJSON struct {
	PeerID         string
	LastHeatBeatAt int64
}

func DecodeSDFSDataNodeInfoJSON(sdfsDataNodeInfoJSON SDFSDataNodeInfoJSON) SDFSDataNodeInfo {
	var ret SDFSDataNodeInfo
	ret.PeerID.SetStr(sdfsDataNodeInfoJSON.PeerID)
	ret.LastHeatBeatAt = time.Unix(sdfsDataNodeInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSDFSDataNodeInfoJSON(sdfsDataNodeInfo SDFSDataNodeInfo) SDFSDataNodeInfoJSON {
	var ret SDFSDataNodeInfoJSON
	ret.PeerID = string(sdfsDataNodeInfo.PeerID.Str())
	ret.LastHeatBeatAt = sdfsDataNodeInfo.LastHeatBeatAt.Unix()
	return ret
}

type SDFSDataNodeInfo struct {
	snettypes.PeerID
	LastHeatBeatAt    time.Time
	LastHeatBeatAtStr string
	SRPCServerAddr    string
	WebServerAddr     string
	sdfsapitypes.DataNodeHeartBeat
}

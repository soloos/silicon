package soloboattypes

import (
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
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
	copy(ret.ID[:], []byte(sdfsDataNodeInfoJSON.PeerID))
	ret.LastHeatBeatAt = time.Unix(sdfsDataNodeInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSDFSDataNodeInfoJSON(sdfsDataNodeInfo SDFSDataNodeInfo) SDFSDataNodeInfoJSON {
	var ret SDFSDataNodeInfoJSON
	ret.PeerID = string(sdfsDataNodeInfo.ID[:])
	ret.LastHeatBeatAt = sdfsDataNodeInfo.LastHeatBeatAt.Unix()
	return ret
}

type SDFSDataNodeInfoUintptr uintptr

func (u SDFSDataNodeInfoUintptr) Ptr() *SDFSDataNodeInfo {
	return (*SDFSDataNodeInfo)(unsafe.Pointer(u))
}

type SDFSDataNodeInfo struct {
	offheap.LKVTableObjectWithBytes64 `db:"-"`
	LastHeatBeatAt                    time.Time
	LastHeatBeatAtStr                 string
	SRPCServerAddr                    string
	WebServerAddr                     string
	sdfsapitypes.DataNodeHeartBeat
}

func (p *SDFSDataNodeInfo) PeerID() snettypes.PeerID { return snettypes.PeerID(p.ID) }

func (p *SDFSDataNodeInfo) PeerIDStr() string { return snettypes.PeerID(p.ID).Str() }

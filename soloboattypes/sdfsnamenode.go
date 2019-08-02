package soloboattypes

import (
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
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
	copy(ret.ID[:], []byte(sdfsNameNodeInfoJSON.PeerID))
	ret.LastHeatBeatAt = time.Unix(sdfsNameNodeInfoJSON.LastHeatBeatAt, 0)
	return ret
}

func EncodeSDFSNameNodeInfoJSON(sdfsNameNodeInfo SDFSNameNodeInfo) SDFSNameNodeInfoJSON {
	var ret SDFSNameNodeInfoJSON
	ret.PeerID = string(sdfsNameNodeInfo.ID[:])
	ret.LastHeatBeatAt = sdfsNameNodeInfo.LastHeatBeatAt.Unix()
	return ret
}

type SDFSNameNodeInfoUintptr uintptr

func (u SDFSNameNodeInfoUintptr) Ptr() *SDFSNameNodeInfo {
	return (*SDFSNameNodeInfo)(unsafe.Pointer(u))
}

type SDFSNameNodeInfo struct {
	offheap.LKVTableObjectWithBytes64 `db:"-"`
	LastHeatBeatAt                    time.Time
	LastHeatBeatAtStr                 string
	SRPCServerAddr                    string
	WebServerAddr                     string
	sdfsapitypes.NameNodeHeartBeat
}

func (p *SDFSNameNodeInfo) PeerID() snettypes.PeerID { return snettypes.PeerID(p.ID) }

func (p *SDFSNameNodeInfo) PeerIDStr() string { return snettypes.PeerID(p.ID).Str() }

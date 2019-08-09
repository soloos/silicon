package soloboatsvr

import (
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SDFSNameNodeDriver) SDFSNameNodeHeartBeat(heartbeat sdfsapitypes.NameNodeHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.SRPCPeerID)
	var uObject, afterSetNewObj = p.sdfsNameNodeTable.MustGetObject(peerID)
	var uSDFSNameNodeInfo = soloboattypes.SDFSNameNodeInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
	}

	uSDFSNameNodeInfo.Ptr().LastHeatBeatAt = time.Now()
	uSDFSNameNodeInfo.Ptr().NameNodeHeartBeat = heartbeat
	err = p.FormatSDFSNameNodeInfo(uSDFSNameNodeInfo.Ptr())
	if err != nil {
		return err
	}

	return nil
}

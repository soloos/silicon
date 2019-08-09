package soloboatsvr

import (
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SDFSDataNodeDriver) SDFSDataNodeHeartBeat(heartbeat sdfsapitypes.DataNodeHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.SRPCPeerID)
	var uObject, afterSetNewObj = p.sdfsDataNodeTable.MustGetObject(peerID)
	var uSDFSDataNodeInfo = soloboattypes.SDFSDataNodeInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
	}

	uSDFSDataNodeInfo.Ptr().LastHeatBeatAt = time.Now()
	uSDFSDataNodeInfo.Ptr().DataNodeHeartBeat = heartbeat
	err = p.FormatSDFSDataNodeInfo(uSDFSDataNodeInfo.Ptr())
	if err != nil {
		return err
	}

	return nil
}

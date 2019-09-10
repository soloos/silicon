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
	var iptr, exists = p.sdfsNameNodeTable.Load(peerID)
	var sdfsNameNodeInfo = soloboattypes.SDFSNameNodeInfo{PeerID: peerID}
	if exists {
		sdfsNameNodeInfo = iptr.(soloboattypes.SDFSNameNodeInfo)
	}

	sdfsNameNodeInfo.LastHeatBeatAt = time.Now()
	sdfsNameNodeInfo.NameNodeHeartBeat = heartbeat
	err = p.FormatSDFSNameNodeInfo(&sdfsNameNodeInfo)
	if err != nil {
		return err
	}

	p.sdfsNameNodeTable.Store(peerID, sdfsNameNodeInfo)

	return nil
}

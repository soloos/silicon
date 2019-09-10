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
	var iptr, exists = p.sdfsDataNodeTable.Load(peerID)
	var sdfsDataNodeInfo = soloboattypes.SDFSDataNodeInfo{PeerID: peerID}
	if exists {
		sdfsDataNodeInfo = iptr.(soloboattypes.SDFSDataNodeInfo)
	}

	sdfsDataNodeInfo.LastHeatBeatAt = time.Now()
	sdfsDataNodeInfo.DataNodeHeartBeat = heartbeat
	err = p.FormatSDFSDataNodeInfo(&sdfsDataNodeInfo)
	if err != nil {
		return err
	}

	p.sdfsDataNodeTable.Store(peerID, sdfsDataNodeInfo)

	return nil
}

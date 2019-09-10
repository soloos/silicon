package soloboatsvr

import (
	"soloos/common/sdbapitypes"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SDBOneDriver) SDBOneHeartBeat(heartbeat sdbapitypes.SDBOneHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.SRPCPeerID)
	var iptr, exists = p.sdbOneTable.Load(peerID)
	var sdbOneInfo = soloboattypes.SDBOneInfo{PeerID: peerID}
	if exists {
		sdbOneInfo = iptr.(soloboattypes.SDBOneInfo)
	}

	sdbOneInfo.LastHeatBeatAt = time.Now()
	sdbOneInfo.SDBOneHeartBeat = heartbeat
	err = p.FormatSDBOneInfo(&sdbOneInfo)
	if err != nil {
		return err
	}

	p.sdbOneTable.Store(peerID, sdbOneInfo)

	return nil
}

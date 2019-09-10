package soloboatsvr

import (
	"soloos/common/snettypes"
	"soloos/common/swalapitypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SWALBrokerDriver) SWALBrokerHeartBeat(heartbeat swalapitypes.BrokerHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.SRPCPeerID)
	var iptr, exists = p.swalBrokerTable.Load(peerID)
	var swalBrokerInfo = soloboattypes.SWALBrokerInfo{PeerID: peerID}
	if exists {
		swalBrokerInfo = iptr.(soloboattypes.SWALBrokerInfo)
	}

	swalBrokerInfo.LastHeatBeatAt = time.Now()
	swalBrokerInfo.BrokerHeartBeat = heartbeat
	err = p.FormatSWALBrokerInfo(&swalBrokerInfo)
	if err != nil {
		return err
	}

	p.swalBrokerTable.Store(peerID, swalBrokerInfo)

	return nil
}

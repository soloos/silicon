package solomq

import (
	"soloos/common/snet"
	"soloos/common/solomqapitypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SolomqDriver) SolomqHeartBeat(heartbeat solomqapitypes.SolomqHeartBeat) error {
	var err error
	var peerID = snet.StrToPeerID(heartbeat.SrpcPeerID)
	var iptr, exists = p.solomqSolomqTable.Load(peerID)
	var solomqSolomqInfo = soloboattypes.SolomqInfo{PeerID: peerID}
	if exists {
		solomqSolomqInfo = iptr.(soloboattypes.SolomqInfo)
	}

	solomqSolomqInfo.LastHeatBeatAt = time.Now()
	solomqSolomqInfo.SolomqHeartBeat = heartbeat
	err = p.FormatSolomqInfo(&solomqSolomqInfo)
	if err != nil {
		return err
	}

	p.solomqSolomqTable.Store(peerID, solomqSolomqInfo)

	return nil
}

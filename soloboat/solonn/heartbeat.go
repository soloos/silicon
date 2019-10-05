package solonn

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/common/solofsapitypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

func (p *SolonnDriver) SolonnHeartBeat(req iron.RequestContext,
	heartbeat solofsapitypes.SolonnHeartBeat,
) iron.Response {
	var (
		peerID       = snettypes.StrToPeerID(heartbeat.SrpcPeerID)
		iptr, exists = p.solonnTable.Load(peerID)
		solonnInfo   = soloboattypes.SolonnInfo{PeerID: peerID}
		err          error
	)

	if exists {
		solonnInfo = iptr.(soloboattypes.SolonnInfo)
	}

	solonnInfo.LastHeatBeatAt = time.Now()
	solonnInfo.SolonnHeartBeat = heartbeat
	err = p.FormatSolonnInfo(&solonnInfo)
	if err != nil {
		return iron.MakeResp(nil, err)
	}

	p.solonnTable.Store(peerID, solonnInfo)

	return iron.MakeResp(nil, nil)
}

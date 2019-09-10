package soloboatsvr

import (
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SDBOneDriver struct {
	soloBoatSvr *SoloBoatSvr
	sdbOneTable sync.Map
}

func (p *SDBOneDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	p.soloBoatSvr = soloBoatSvr
	return nil
}

func (p *SDBOneDriver) FormatSDBOneInfo(sdbOneInfo *soloboattypes.SDBOneInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	sdbOneInfo.LastHeatBeatAtStr = sdbOneInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(sdbOneInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	sdbOneInfo.SRPCServerAddr = peer.AddressStr()

	//TODO enable WebServer
	peer, _ = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(sdbOneInfo.WebPeerID))
	sdbOneInfo.WebServerAddr = peer.AddressStr()

	return nil
}

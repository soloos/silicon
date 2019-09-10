package soloboatsvr

import (
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SWALBrokerDriver struct {
	soloBoatSvr     *SoloBoatSvr
	swalBrokerTable sync.Map
}

func (p *SWALBrokerDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	p.soloBoatSvr = soloBoatSvr
	return nil
}

func (p *SWALBrokerDriver) FormatSWALBrokerInfo(swalBrokerInfo *soloboattypes.SWALBrokerInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	swalBrokerInfo.LastHeatBeatAtStr = swalBrokerInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(swalBrokerInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	swalBrokerInfo.SRPCServerAddr = peer.AddressStr()

	//TODO enable WebServer
	peer, _ = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(swalBrokerInfo.WebPeerID))
	swalBrokerInfo.WebServerAddr = peer.AddressStr()

	return nil
}

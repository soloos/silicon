package solodb

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SolodbDriver struct {
	soloboatIns *soloboattypes.Soloboat
	solodbTable sync.Map
}

var _ = iron.IServer(&SolodbDriver{})

func (p *SolodbDriver) ServerName() string {
	return "Soloos.Soloboat.SolodbDriver"
}

func (p *SolodbDriver) Init(soloboatIns *soloboattypes.Soloboat) error {
	p.soloboatIns = soloboatIns
	return nil
}

func (p *SolodbDriver) Serve() error {
	return nil
}

func (p *SolodbDriver) Close() error {
	return nil
}

func (p *SolodbDriver) FormatSolodbInfo(solodbInfo *soloboattypes.SolodbInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	solodbInfo.LastHeatBeatAtStr = solodbInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloboatIns.SNetDriver.GetPeer(snettypes.StrToPeerID(solodbInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	solodbInfo.SRPCServerAddr = peer.AddressStr()

	//TODO enable WebServer
	peer, _ = p.soloboatIns.SNetDriver.GetPeer(snettypes.StrToPeerID(solodbInfo.WebPeerID))
	solodbInfo.WebServerAddr = peer.AddressStr()

	return nil
}

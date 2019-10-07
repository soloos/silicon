package solodb

import (
	"soloos/common/iron"
	"soloos/common/snet"
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SolodbDriver struct {
	soloboatIns soloboattypes.Soloboat
	*soloosbase.SoloosEnv
	solodbTable sync.Map
}

var _ = iron.IServer(&SolodbDriver{})

func (p *SolodbDriver) ServerName() string {
	return "Soloos.Soloboat.SolodbDriver"
}

func (p *SolodbDriver) Init(soloboatIns soloboattypes.Soloboat) error {
	p.soloboatIns = soloboatIns
	p.SoloosEnv = p.soloboatIns.GetSoloosEnv()
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
		peer snet.Peer
		err  error
	)

	solodbInfo.LastHeatBeatAtStr = solodbInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.SNetDriver.GetPeer(snet.StrToPeerID(solodbInfo.SrpcPeerID))
	if err != nil {
		return err
	}
	solodbInfo.SrpcServerAddr = peer.AddressStr()

	//TODO enable WebServer
	peer, _ = p.SNetDriver.GetPeer(snet.StrToPeerID(solodbInfo.WebPeerID))
	solodbInfo.WebServerAddr = peer.AddressStr()

	return nil
}

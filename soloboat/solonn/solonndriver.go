package solonn

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SolonnDriver struct {
	soloboat soloboattypes.Soloboat
	*soloosbase.SoloosEnv
	solonnTable sync.Map
}

var _ = iron.IServer(&SolonnDriver{})

func (p *SolonnDriver) ServerName() string {
	return "Soloos.Soloboat.SolonnDriver"
}

func (p *SolonnDriver) Init(soloboatIns soloboattypes.Soloboat) error {
	var err error
	p.soloboat = soloboatIns
	p.SoloosEnv = p.soloboat.GetSoloosEnv()

	err = p.initService()
	if err != nil {
		return err
	}

	return nil
}

func (p *SolonnDriver) Serve() error {
	return nil
}

func (p *SolonnDriver) Close() error {
	return nil
}

func (p *SolonnDriver) FormatSolonnInfo(solonnInfo *soloboattypes.SolonnInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	solonnInfo.LastHeatBeatAtStr = solonnInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.SNetDriver.GetPeer(snettypes.StrToPeerID(solonnInfo.SrpcPeerID))
	if err != nil {
		return err
	}
	solonnInfo.SrpcServerAddr = peer.AddressStr()

	peer, err = p.SNetDriver.GetPeer(snettypes.StrToPeerID(solonnInfo.WebPeerID))
	if err != nil {
		return err
	}
	solonnInfo.WebServerAddr = peer.AddressStr()

	return nil
}

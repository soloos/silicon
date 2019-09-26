package solodn

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SolodnDriver struct {
	soloboatIns *soloboattypes.Soloboat
	solodnTable sync.Map
}

var _ = iron.IServer(&SolodnDriver{})

func (p *SolodnDriver) ServerName() string {
	return "Soloos.Soloboat.SolodnDriver"
}

func (p *SolodnDriver) Init(soloboatIns *soloboattypes.Soloboat) error {
	p.soloboatIns = soloboatIns
	return nil
}

func (p *SolodnDriver) Serve() error {
	return nil
}

func (p *SolodnDriver) Close() error {
	return nil
}

func (p *SolodnDriver) FormatSolodnInfo(solodnInfo *soloboattypes.SolodnInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	solodnInfo.LastHeatBeatAtStr = solodnInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloboatIns.SNetDriver.GetPeer(snettypes.StrToPeerID(solodnInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	solodnInfo.SRPCServerAddr = peer.AddressStr()

	peer, err = p.soloboatIns.SNetDriver.GetPeer(snettypes.StrToPeerID(solodnInfo.WebPeerID))
	if err != nil {
		return err
	}
	solodnInfo.WebServerAddr = peer.AddressStr()

	return nil
}

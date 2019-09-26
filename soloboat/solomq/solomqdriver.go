package solomq

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SolomqDriver struct {
	soloboatIns       *soloboattypes.Soloboat
	solomqSolomqTable sync.Map
}

var _ = iron.IServer(&SolomqDriver{})

func (p *SolomqDriver) ServerName() string {
	return "Soloos.Soloboat.SolomqDriver"
}

func (p *SolomqDriver) Init(soloboatIns *soloboattypes.Soloboat) error {
	p.soloboatIns = soloboatIns
	return nil
}

func (p *SolomqDriver) Serve() error {
	return nil
}

func (p *SolomqDriver) Close() error {
	return nil
}

func (p *SolomqDriver) FormatSolomqInfo(solomqSolomqInfo *soloboattypes.SolomqInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	solomqSolomqInfo.LastHeatBeatAtStr = solomqSolomqInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloboatIns.SNetDriver.GetPeer(snettypes.StrToPeerID(solomqSolomqInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	solomqSolomqInfo.SRPCServerAddr = peer.AddressStr()

	//TODO enable WebServer
	peer, _ = p.soloboatIns.SNetDriver.GetPeer(snettypes.StrToPeerID(solomqSolomqInfo.WebPeerID))
	solomqSolomqInfo.WebServerAddr = peer.AddressStr()

	return nil
}

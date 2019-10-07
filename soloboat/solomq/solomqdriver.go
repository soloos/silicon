package solomq

import (
	"soloos/common/iron"
	"soloos/common/snet"
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SolomqDriver struct {
	soloboatIns soloboattypes.Soloboat
	*soloosbase.SoloosEnv
	solomqSolomqTable sync.Map
}

var _ = iron.IServer(&SolomqDriver{})

func (p *SolomqDriver) ServerName() string {
	return "Soloos.Soloboat.SolomqDriver"
}

func (p *SolomqDriver) Init(soloboatIns soloboattypes.Soloboat) error {
	p.soloboatIns = soloboatIns
	p.SoloosEnv = p.soloboatIns.GetSoloosEnv()
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
		peer snet.Peer
		err  error
	)

	solomqSolomqInfo.LastHeatBeatAtStr = solomqSolomqInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.SNetDriver.GetPeer(snet.StrToPeerID(solomqSolomqInfo.SrpcPeerID))
	if err != nil {
		return err
	}
	solomqSolomqInfo.SrpcServerAddr = peer.AddressStr()

	//TODO enable WebServer
	peer, _ = p.SNetDriver.GetPeer(snet.StrToPeerID(solomqSolomqInfo.WebPeerID))
	solomqSolomqInfo.WebServerAddr = peer.AddressStr()

	return nil
}

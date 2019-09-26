package soloboat

import (
	"soloos/common/iron"
	"soloos/common/log"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
)

type SNetDriver struct {
	soloboatIns *soloboattypes.Soloboat
}

var _ = iron.IServer(&SNetDriver{})

func (p *SNetDriver) Init(soloboatIns *soloboattypes.Soloboat) error {
	var err error
	p.soloboatIns = soloboatIns

	err = p.soloboatIns.SoloOSEnv.SNetDriver.PrepareServer("/Api/SNet", &p.soloboatIns.webServer.server,
		p.FetchSNetPeerFromDB,
		p.RegisterSNetPeerInDB)
	if err != nil {
		return err
	}

	// register myself
	err = p.soloboatIns.SoloOSEnv.SNetDriver.RegisterPeerInDB(p.soloboatIns.webPeer)
	if err != nil {
		return err
	}

	return nil
}

func (p *SNetDriver) ServerName() string {
	return "soloboattypes.Soloboat.SNetDriver"
}

func (p *SNetDriver) Serve() error {
	var err error

	err = p.ListSNetPeerFromDB(func(peer snettypes.Peer) bool {
		var err = p.soloboatIns.SoloOSEnv.SNetDriver.RegisterPeer(peer)
		if err != nil {
			log.Error("RegisterPeer error, err:", err)
			return false
		}
		return true
	})
	if err != nil {
		return err
	}

	err = p.soloboatIns.SoloOSEnv.SNetDriver.ServerServe()
	if err != nil {
		return err
	}

	return nil
}

func (p *SNetDriver) Close() error {
	return p.soloboatIns.SoloOSEnv.SNetDriver.CloseServer()
}

package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/log"
	"soloos/common/snettypes"
)

type SNetDriver struct {
	soloBoatSvr *SoloBoatSvr
}

var _ = iron.IServer(&SNetDriver{})

func (p *SNetDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error
	p.soloBoatSvr = soloBoatSvr

	err = p.soloBoatSvr.SoloOSEnv.SNetDriver.PrepareServer("/Api/SNet", &p.soloBoatSvr.webServer.server,
		p.FetchSNetPeerFromDB,
		p.RegisterSNetPeerInDB)
	if err != nil {
		return err
	}

	// register myself
	err = p.soloBoatSvr.SoloOSEnv.SNetDriver.RegisterPeerInDB(p.soloBoatSvr.webPeer)
	if err != nil {
		return err
	}

	return nil
}

func (p *SNetDriver) ServerName() string {
	return "SoloBoatSvr.SNetDriver"
}

func (p *SNetDriver) Serve() error {
	var err error

	err = p.ListSNetPeerFromDB(func(peer snettypes.Peer) bool {
		var err = p.soloBoatSvr.SoloOSEnv.SNetDriver.RegisterPeer(peer)
		if err != nil {
			log.Error("RegisterPeer error, err:", err)
			return false
		}
		return true
	})
	if err != nil {
		return err
	}

	err = p.soloBoatSvr.SoloOSEnv.SNetDriver.ServerServe()
	if err != nil {
		return err
	}

	return nil
}

func (p *SNetDriver) Close() error {
	return p.soloBoatSvr.SoloOSEnv.SNetDriver.CloseServer()
}

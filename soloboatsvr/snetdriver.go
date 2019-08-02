package soloboatsvr

import (
	"soloos/common/log"
	"soloos/common/snettypes"
)

type SNetDriver struct {
	soloBoatSvr *SoloBoatSvr
}

func (p *SNetDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error
	p.soloBoatSvr = soloBoatSvr

	err = p.soloBoatSvr.SoloOSEnv.SNetDriver.PrepareServer(p.soloBoatSvr.options.SNetDriverListenAddr,
		p.soloBoatSvr.options.SNetDriverServeAddr,
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

package soloboatsvr

import (
	"soloos/common/snettypes"
)

type SNetDriver struct {
	soloBoatSvr *SoloBoatSvr
}

func (p *SNetDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	p.soloBoatSvr = soloBoatSvr
	return nil
}

func (p *SNetDriver) Serve() error {
	var err error

	err = p.ListSNetPeerFromDB(func(peer snettypes.Peer) bool {
		err = p.soloBoatSvr.SoloOSEnv.SNetDriver.RegisterPeer(peer)
		if err != nil {
			return false
		}
		return true
	})
	if err != nil {
		return err
	}

	err = p.soloBoatSvr.SoloOSEnv.SNetDriver.StartServer(p.soloBoatSvr.options.SNetDriverListenAddr,
		p.soloBoatSvr.options.SNetDriverServeAddr,
		p.FetchSNetPeerFromDB,
		p.RegisterSNetPeerInDB)
	if err != nil {
		return err
	}

	return nil
}

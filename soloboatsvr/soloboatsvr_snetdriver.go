package soloboatsvr

import (
	"soloos/common/snettypes"
)

func (p *SoloBoatSvr) StartSNetDriverServer() error {
	var err error

	err = p.ListSNetPeerFromDB(func(peer snettypes.Peer) bool {
		err = p.SoloOSEnv.SNetDriver.RegisterPeer(peer)
		if err != nil {
			return false
		}
		return true
	})
	if err != nil {
		return err
	}

	err = p.SoloOSEnv.SNetDriver.StartServer(p.options.SNetDriverListenAddr,
		p.options.SNetDriverServeAddr,
		p.FetchSNetPeerFromDB,
		p.RegisterSNetPeerInDB)
	if err != nil {
		return err
	}

	return nil
}

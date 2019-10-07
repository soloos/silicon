package soloboat

import (
	"soloos/common/iron"
	"soloos/common/log"
	"soloos/common/snet"
	"soloos/common/solodbapi"
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboattypes"
)

type SNetDriver struct {
	soloboatIns soloboattypes.Soloboat
	*soloosbase.SoloosEnv
	dbConn *solodbapi.Connection
}

var _ = iron.IServer(&SNetDriver{})

func (p *SNetDriver) Init(soloboatIns soloboattypes.Soloboat) error {
	var err error
	p.soloboatIns = soloboatIns
	p.SoloosEnv = soloboatIns.GetSoloosEnv()
	p.dbConn = p.soloboatIns.GetDBConn()

	err = p.SoloosEnv.SNetDriver.PrepareServer("/Api/SNet", p.soloboatIns.GetWebServer(),
		p.FetchSNetPeerFromDB,
		p.RegisterSNetPeerInDB)
	if err != nil {
		return err
	}

	// register myself
	err = p.SoloosEnv.SNetDriver.RegisterPeerInDB(p.soloboatIns.GetWebPeer())
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

	err = p.ListSNetPeerFromDB(func(peer snet.Peer) bool {
		var err = p.SoloosEnv.SNetDriver.RegisterPeer(peer)
		if err != nil {
			log.Error("RegisterPeer error, err:", err)
			return false
		}
		return true
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *SNetDriver) Close() error {
	return p.SoloosEnv.SNetDriver.CloseServer()
}

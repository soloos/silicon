package soloboat

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/common/solodbapi"
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboattypes"
)

type Soloboat struct {
	*soloosbase.SoloOSEnv
	options SoloboatOptions
	webPeer snettypes.Peer
	dbConn  solodbapi.Connection

	webServer  WebServer
	snetDriver SNetDriver
	SoloboatDrivers
	serverDriver iron.ServerDriver
}

func (p *Soloboat) initSNetPeer() error {
	p.webPeer.ID = snettypes.StrToPeerID(p.options.WebPeerID)
	p.webPeer.SetAddress(p.options.WebServerOptions.ServeStr)
	p.webPeer.ServiceProtocol = soloboattypes.DefaultSoloboatRPCProtocol

	return nil
}

func (p *Soloboat) initDBConn() error {
	var err error
	err = p.dbConn.Init(p.options.DBDriver, p.options.Dsn)
	if err != nil {
		return err
	}

	err = p.installSchema(p.options.DBDriver)
	if err != nil {
		return err
	}

	return nil
}

func (p *Soloboat) initSidecarDriver() error {
	return nil
}

func (p *Soloboat) Init(soloOSEnv *soloosbase.SoloOSEnv, options SoloboatOptions) error {
	var err error

	p.SoloOSEnv = soloOSEnv
	p.options = options

	err = p.preparePProf(p.options.PProfListenAddr)
	if err != nil {
		return err
	}

	err = p.initSNetPeer()
	if err != nil {
		return err
	}

	err = p.initDBConn()
	if err != nil {
		return err
	}

	err = p.webServer.Init(p)
	if err != nil {
		return err
	}

	err = p.SoloboatDrivers.Init(p)
	if err != nil {
		return err
	}

	err = p.serverDriver.Init(&p.webServer, &p.SoloboatDrivers)
	if err != nil {
		return err
	}

	return nil
}

func (p *Soloboat) GetPeerID() snettypes.PeerID {
	return p.webPeer.ID
}

func (p *Soloboat) Serve() error {
	return p.serverDriver.Serve()
}

func (p *Soloboat) Close() error {
	return p.serverDriver.Close()
}

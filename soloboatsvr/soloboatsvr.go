package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/common/sdbapi"
	"soloos/common/snettypes"
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboattypes"
)

type SoloBoatSvr struct {
	*soloosbase.SoloOSEnv
	options SoloBoatSvrOptions
	webPeer snettypes.Peer
	dbConn  sdbapi.Connection

	webServer          WebServer
	snetDriver         SNetDriver
	sideCarDriver      SideCarDriver
	sdfsNameNodeDriver SDFSNameNodeDriver
	sdfsDataNodeDriver SDFSDataNodeDriver
	swalBrokerDriver   SWALBrokerDriver
	sdbOneDriver       SDBOneDriver

	serverDriver iron.ServerDriver
}

func (p *SoloBoatSvr) initSNetPeer() error {
	p.webPeer.ID = snettypes.StrToPeerID(p.options.WebPeerID)
	p.webPeer.SetAddress(p.options.WebServerOptions.ServeStr)
	p.webPeer.ServiceProtocol = soloboattypes.DefaultSoloBoatRPCProtocol

	return nil
}

func (p *SoloBoatSvr) initDBConn() error {
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

func (p *SoloBoatSvr) initSideCarDriver() error {
	return nil
}

func (p *SoloBoatSvr) Init(soloOSEnv *soloosbase.SoloOSEnv, options SoloBoatSvrOptions) error {
	var err error

	p.SoloOSEnv = soloOSEnv
	p.options = options

	err = p.initSNetPeer()
	if err != nil {
		return err
	}

	err = p.initDBConn()
	if err != nil {
		return err
	}

	err = p.snetDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.webServer.Init(p)
	if err != nil {
		return err
	}

	err = p.sideCarDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.sdfsNameNodeDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.sdfsDataNodeDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.swalBrokerDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.sdbOneDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.serverDriver.Init(&p.webServer, &p.snetDriver)
	if err != nil {
		return err
	}

	return nil
}

func (p *SoloBoatSvr) GetPeerID() snettypes.PeerID {
	return p.webPeer.ID
}

func (p *SoloBoatSvr) Serve() error {
	return p.serverDriver.Serve()
}

func (p *SoloBoatSvr) Close() error {
	return p.serverDriver.Close()
}

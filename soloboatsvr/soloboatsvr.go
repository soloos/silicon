package soloboatsvr

import (
	"soloos/common/log"
	"soloos/common/sdbapi"
	"soloos/common/snettypes"
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboattypes"
)

type SoloBoatSvr struct {
	*soloosbase.SoloOSEnv
	options       SoloBoatSvrOptions
	webPeer       snettypes.Peer
	dbConn        sdbapi.Connection
	webServer     WebServer
	servicesCount int

	snetDriver         SNetDriver
	sideCarDriver      SideCarDriver
	sdfsNameNodeDriver SDFSNameNodeDriver
	sdfsDataNodeDriver SDFSDataNodeDriver
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

	p.servicesCount = 2

	return nil
}

func (p *SoloBoatSvr) GetPeerID() snettypes.PeerID {
	return p.webPeer.ID
}

func (p *SoloBoatSvr) Serve() error {
	var errChans = make(chan error, p.servicesCount)

	go func(errChans chan error) {
		errChans <- p.snetDriver.Serve()
	}(errChans)

	go func(errChans chan error) {
		errChans <- p.webServer.Serve()
	}(errChans)

	var err error
	for i := 0; i < p.servicesCount; i++ {
		var tmperr = <-errChans
		if tmperr != nil {
			log.Error("SoloBoatSvr Serve error, err:", tmperr)
			err = tmperr
		}
	}

	return err
}

func (p *SoloBoatSvr) Close() error {
	var err error
	err = p.SoloOSEnv.SNetDriver.CloseServer()
	if err != nil {
		return err
	}

	err = p.webServer.Close()
	if err != nil {
		return err
	}

	return nil
}

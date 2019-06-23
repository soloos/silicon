package agent

import (
	"soloos/common/iron"
	"soloos/common/sdbapi"
	"soloos/common/siliconapitypes"
	"soloos/common/snettypes"
	"soloos/common/soloosbase"
)

type SiliconAgent struct {
	*soloosbase.SoloOSEnv
	options   SiliconAgentOptions
	peer      snettypes.Peer
	dbConn    sdbapi.Connection
	webServer iron.Server
}

func (p *SiliconAgent) initSNetPeer() error {
	var err error

	p.peer.ID = snettypes.StrToPeerID(p.options.PeerID)
	p.peer.SetAddress(p.options.WebServerOptions.ServeStr)
	p.peer.ServiceProtocol = siliconapitypes.DefaultSiliconRPCProtocol

	err = p.SoloOSEnv.SNetDriver.RegisterPeer(p.peer)
	if err != nil {
		return err
	}

	p.SoloOSEnv.SNetDriver.StartServer(p.options.SNetDriverListenAddr,
		p.options.SNetDriverServeAddr,
		nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *SiliconAgent) initDBConn() error {
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

func (p *SiliconAgent) Init(soloOSEnv *soloosbase.SoloOSEnv, options SiliconAgentOptions) error {
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

	return nil
}

func (p *SiliconAgent) GetPeerID() snettypes.PeerID {
	return p.peer.ID
}

func (p *SiliconAgent) Serve() error {
	var err error
	err = p.webServer.Serve()
	return err
}

func (p *SiliconAgent) Close() error {
	var err error
	err = p.webServer.Close()
	if err != nil {
		return err
	}

	return nil
}

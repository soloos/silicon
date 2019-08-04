package soloboatsvr

import (
	"soloos/common/iron"
)

type WebServer struct {
	soloBoatSvr *SoloBoatSvr
	server      iron.Server
}

var _ = iron.IServer(&WebServer{})

func (p *WebServer) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error

	p.soloBoatSvr = soloBoatSvr
	err = p.server.Init(p.soloBoatSvr.options.WebServerOptions)
	if err != nil {
		return err
	}

	err = p.prepareCtr()
	if err != nil {
		return err
	}

	err = p.prepareApi()
	if err != nil {
		return err
	}

	return nil
}

func (p *WebServer) ServerName() string {
	return "SoloBoatSvr.WebServer"
}

func (p *WebServer) Serve() error {
	var err error
	err = p.server.Serve()
	if err != nil {
		return err
	}

	return nil
}

func (p *WebServer) Close() error {
	var err error
	err = p.server.Close()
	if err != nil {
		return err
	}

	return nil
}

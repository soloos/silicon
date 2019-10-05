package soloboat

import (
	"soloos/common/iron"
)

func (p *Soloboat) GetWebServer() *iron.Server {
	return &p.webServer.server
}

type WebServer struct {
	soloboatIns *Soloboat
	server      iron.Server
}

var _ = iron.IServer(&WebServer{})

func (p *WebServer) ServerName() string {
	return "Soloboat.WebServer"
}

func (p *WebServer) Init(soloboatIns *Soloboat) error {
	var err error

	p.soloboatIns = soloboatIns
	err = p.server.Init(p.soloboatIns.options.WebServerOptions)
	if err != nil {
		return err
	}

	err = p.prepareApi()
	if err != nil {
		return err
	}

	return nil
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

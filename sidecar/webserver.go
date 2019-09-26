package sidecar

import "soloos/common/iron"

type WebServer struct {
	sideCar    *Sidecar
	httpServer iron.Server
}

func (p *WebServer) Init(sideCar *Sidecar, options iron.Options) error {
	var err error
	p.sideCar = sideCar

	err = p.httpServer.Init(options)
	if err != nil {
		return err
	}

	return nil
}

func (p *WebServer) Serve() error {
	return p.httpServer.Serve()
}

func (p *WebServer) Close() error {
	return p.httpServer.Close()
}

package agent

import (
	"soloos/common/iron"
)

type WebServer struct {
	agent  *SiliconAgent
	server iron.Server
}

func (p *WebServer) Init(agent *SiliconAgent) error {
	var err error

	p.agent = agent
	err = p.server.Init(p.agent.options.WebServerOptions)
	if err != nil {
		return err
	}

	p.server.Router("/", p.ctrMain)

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

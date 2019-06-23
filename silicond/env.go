package main

import (
	"soloos/common/soloosbase"
	"soloos/silicon/agent"
)

type Env struct {
	SoloOSEnv soloosbase.SoloOSEnv

	SiliconAgent agent.SiliconAgent
}

func (p *Env) Init(options Options) error {
	var err error

	err = p.SoloOSEnv.Init()
	if err != nil {
		return err
	}

	err = p.SiliconAgent.Init(&p.SoloOSEnv, options.SiliconAgentOptions)
	if err != nil {
		return err
	}

	return nil
}

func (p *Env) Serve() error {
	return p.SiliconAgent.Serve()
}

func (p *Env) Close() error {
	return nil
}

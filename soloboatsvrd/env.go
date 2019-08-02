package main

import (
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboatsvr"
)

type Env struct {
	SoloOSEnv soloosbase.SoloOSEnv

	SoloBoatSvr soloboatsvr.SoloBoatSvr
}

func (p *Env) Init(options Options) error {
	var err error

	err = p.SoloOSEnv.InitWithSNet("")
	if err != nil {
		return err
	}

	err = p.SoloBoatSvr.Init(&p.SoloOSEnv, options.SoloBoatSvrOptions)
	if err != nil {
		return err
	}

	return nil
}

func (p *Env) Serve() error {
	return p.SoloBoatSvr.Serve()
}

func (p *Env) Close() error {
	return nil
}

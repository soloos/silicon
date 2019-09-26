package soloboatd

import (
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboat"
)

type SoloboatDaemon struct {
	SoloOSEnv   soloosbase.SoloOSEnv
	Soloboat soloboat.Soloboat
}

func (p *SoloboatDaemon) Init(options Options) error {
	var err error

	err = p.SoloOSEnv.InitWithSNet("")
	if err != nil {
		return err
	}

	err = p.Soloboat.Init(&p.SoloOSEnv, options.SoloboatOptions)
	if err != nil {
		return err
	}

	return nil
}

func (p *SoloboatDaemon) Serve() error {
	return p.Soloboat.Serve()
}

func (p *SoloboatDaemon) Close() error {
	return nil
}

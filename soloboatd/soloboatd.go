package soloboatd

import (
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboat"
)

type SoloboatDaemon struct {
	SoloosEnv   soloosbase.SoloosEnv
	Soloboat soloboat.Soloboat
}

func (p *SoloboatDaemon) Init(options Options) error {
	var err error

	err = p.SoloosEnv.InitWithSNet("")
	if err != nil {
		return err
	}

	err = p.Soloboat.Init(&p.SoloosEnv, options.SoloboatOptions)
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

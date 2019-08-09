package soloboatsvrd

import (
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboatsvr"
)

type SoloBoatSvrd struct {
	SoloOSEnv   soloosbase.SoloOSEnv
	SoloBoatSvr soloboatsvr.SoloBoatSvr
}

func (p *SoloBoatSvrd) Init(options Options) error {
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

func (p *SoloBoatSvrd) Serve() error {
	return p.SoloBoatSvr.Serve()
}

func (p *SoloBoatSvrd) Close() error {
	return nil
}

package solofs

import (
	"soloos/common/iron"
	"soloos/common/soloosbase"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SolofsDriver struct {
	soloboatIns soloboattypes.Soloboat
	*soloosbase.SoloosEnv
	solofsTable sync.Map
}

var _ = iron.IServer(&SolofsDriver{})

func (p *SolofsDriver) ServerName() string {
	return "Soloos.Soloboat.SolofsDriver"
}

func (p *SolofsDriver) Init(soloboatIns soloboattypes.Soloboat) error {
	p.soloboatIns = soloboatIns
	p.SoloosEnv = p.soloboatIns.GetSoloosEnv()
	return nil
}

func (p *SolofsDriver) Serve() error {
	return nil
}

func (p *SolofsDriver) Close() error {
	return nil
}

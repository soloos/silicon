package soloboat

import (
	"soloos/common/iron"
	"soloos/soloboat/soloboat/sidecar"
	"soloos/soloboat/soloboat/solodb"
	"soloos/soloboat/soloboat/solodn"
	"soloos/soloboat/soloboat/solomq"
	"soloos/soloboat/soloboat/solonn"
)

type SoloboatDrivers struct {
	soloboat           *Soloboat
	sideCarDriver      sidecar.SidecarDriver
	solonnDriver       solonn.SolonnDriver
	solodnDriver       solodn.SolodnDriver
	solomqSolomqDriver solomq.SolomqDriver
	solodbDriver       solodb.SolodbDriver

	serverDriver iron.ServerDriver
}

var _ = iron.IServer(&SoloboatDrivers{})

func (p *SoloboatDrivers) ServerName() string {
	return "Soloos.Soloboat.Drivers"
}

func (p *SoloboatDrivers) Init(soloboatIns *Soloboat) error {
	var err error
	p.soloboat = soloboatIns

	err = p.serverDriver.Init()
	if err != nil {
		return err
	}

	err = p.snetDriver.Init(p.soloboat)
	if err != nil {
		return err
	}
	p.serverDriver.AddServer(&p.snetDriver)

	err = p.sideCarDriver.Init(p.soloboat)
	if err != nil {
		return err
	}
	p.serverDriver.AddServer(&p.sideCarDriver)

	err = p.solonnDriver.Init(p.soloboat)
	if err != nil {
		return err
	}
	p.serverDriver.AddServer(&p.solonnDriver)

	err = p.solodnDriver.Init(p.soloboat)
	if err != nil {
		return err
	}
	p.serverDriver.AddServer(&p.solodnDriver)

	err = p.solomqSolomqDriver.Init(p.soloboat)
	if err != nil {
		return err
	}
	p.serverDriver.AddServer(&p.solomqSolomqDriver)

	err = p.solodbDriver.Init(p.soloboat)
	if err != nil {
		return err
	}
	p.serverDriver.AddServer(&p.solodbDriver)

	return nil
}

func (p *SoloboatDrivers) Serve() error {
	return p.serverDriver.Serve()
}

func (p *SoloboatDrivers) Close() error {
	return p.serverDriver.Close()
}

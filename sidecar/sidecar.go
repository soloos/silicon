package sidecar

import (
	"soloos/common/util"
)

type SideCar struct {
	SDFSDriver   SDFSDriver
	SWALDriver   SWALDriver
	BadgerDriver BadgerDriver

	WebServer WebServer
}

func (p *SideCar) Init(options Options) error {
	var err error
	err = p.SDFSDriver.Init(p, options.SDFS)
	if err != nil {
		return err
	}

	err = p.SWALDriver.Init(p, options.SWAL)
	if err != nil {
		return err
	}

	err = p.BadgerDriver.Init(p, options.Badger)
	if err != nil {
		return err
	}

	err = p.WebServer.Init(p, options.WebServer)
	if err != nil {
		return err
	}

	return nil
}

func (p *SideCar) Serve() error {
	go func() {
		util.AssertErrIsNil(p.WebServer.Serve())
	}()
	return nil
}

func (p *SideCar) Close() error {
	return p.WebServer.Close()
}

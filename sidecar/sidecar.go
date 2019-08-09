package sidecar

import (
	"soloos/common/snettypes"
	"soloos/common/soloosbase"
	"soloos/common/util"
)

type SideCar struct {
	*soloosbase.SoloOSEnv
	SDFSDriver   SDFSDriver
	SWALDriver   SWALDriver
	BadgerDriver BadgerDriver

	srpcPeer snettypes.Peer
	webPeer  snettypes.Peer

	WebServer WebServer

	heartBeatServerOptionsArr []snettypes.HeartBeatServerOptions
}

func (p *SideCar) Init(soloOSEnv *soloosbase.SoloOSEnv, options Options) error {
	var err error
	err = p.SDFSDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.SWALDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.BadgerDriver.Init(p)
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

package sidecar

import (
	"soloos/common/snettypes"
	"soloos/common/soloosbase"
	"soloos/common/util"
)

type Sidecar struct {
	*soloosbase.SoloOSEnv
	SolofsDriver   SolofsDriver
	SolomqDriver   SolomqDriver
	BadgerDriver BadgerDriver

	srpcPeer snettypes.Peer
	webPeer  snettypes.Peer

	WebServer WebServer

	heartBeatServerOptionsArr []snettypes.HeartBeatServerOptions
}

func (p *Sidecar) Init(soloOSEnv *soloosbase.SoloOSEnv, options Options) error {
	var err error
	err = p.SolofsDriver.Init(p)
	if err != nil {
		return err
	}

	err = p.SolomqDriver.Init(p)
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

func (p *Sidecar) Serve() error {
	go func() {
		util.AssertErrIsNil(p.WebServer.Serve())
	}()
	return nil
}

func (p *Sidecar) Close() error {
	return p.WebServer.Close()
}

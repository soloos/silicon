package sidecar

import (
	"soloos/common/snet"
	"soloos/common/soloosbase"
	"soloos/common/util"
)

type Sidecar struct {
	*soloosbase.SoloosEnv
	SolofsDriver SolofsDriver
	SolomqDriver SolomqDriver
	BadgerDriver BadgerDriver

	srpcPeer snet.Peer
	webPeer  snet.Peer

	WebServer WebServer

	heartBeatServerOptionsArr []snet.HeartBeatServerOptions
}

func (p *Sidecar) Init(soloosEnv *soloosbase.SoloosEnv, options Options) error {
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

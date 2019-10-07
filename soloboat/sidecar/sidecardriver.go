package sidecar

import (
	"soloos/common/iron"
	"soloos/common/snet"
	"soloos/common/solodbapi"
	"soloos/common/soloosbase"
	"soloos/soloboat/sidecartypes"
	"soloos/soloboat/soloboattypes"
	"soloos/solodb/offheap"
	"time"
)

type SidecarDriver struct {
	soloboatIns soloboattypes.Soloboat
	*soloosbase.SoloosEnv
	dbConn       *solodbapi.Connection
	SidecarTable offheap.LKVTableWithBytes64
}

var _ = iron.IServer(&SidecarDriver{})

func (p *SidecarDriver) ServerName() string {
	return "Soloos.Soloboat.SidecarDriver"
}

func (p *SidecarDriver) Init(soloboatIns soloboattypes.Soloboat) error {
	var err error
	p.soloboatIns = soloboatIns
	p.SoloosEnv = p.soloboatIns.GetSoloosEnv()
	p.dbConn = p.soloboatIns.GetDBConn()
	err = p.SoloosEnv.OffheapDriver.InitLKVTableWithBytes64(&p.SidecarTable, "Sidecar",
		int(soloboattypes.SidecarInfoStructSize), -1, offheap.DefaultKVTableSharedCount, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *SidecarDriver) Serve() error {
	return nil
}

func (p *SidecarDriver) Close() error {
	return nil
}

func (p *SidecarDriver) SidecarHeartBeat(heartbeat sidecartypes.SidecarHeartBeat) error {
	var err error
	var peerID = snet.StrToPeerID(heartbeat.WebPeerID)
	var uObject, afterSetNewObj = p.SidecarTable.MustGetObject(peerID)
	var uSidecarInfo = soloboattypes.SidecarInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
	}

	uSidecarInfo.Ptr().LastHeatBeatAt = time.Now()
	uSidecarInfo.Ptr().SidecarHeartBeat = heartbeat
	err = p.FormatSidecarInfo(uSidecarInfo.Ptr())
	if err != nil {
		return err
	}

	return nil
}

func (p *SidecarDriver) FormatSidecarInfo(SidecarInfo *soloboattypes.SidecarInfo) error {
	var (
		peer snet.Peer
		err  error
	)

	peer, err = p.SNetDriver.GetPeer(snet.StrToPeerID(SidecarInfo.WebPeerID))
	if err != nil {
		return err
	}
	SidecarInfo.WebServerAddr = peer.AddressStr()

	return nil
}

func (p *SidecarDriver) ListObject(listPeer offheap.LKVTableListObjectWithBytes64) {
	p.SidecarTable.ListObject(listPeer)
}

package soloboatsvr

import (
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"soloos/soloboat/sidecartypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

type SideCarDriver struct {
	soloBoatSvr  *SoloBoatSvr
	SideCarTable offheap.LKVTableWithBytes64
}

func (p *SideCarDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error
	p.soloBoatSvr = soloBoatSvr
	err = p.soloBoatSvr.SoloOSEnv.OffheapDriver.InitLKVTableWithBytes64(&p.SideCarTable, "SideCar",
		int(soloboattypes.SideCarInfoStructSize), -1, offheap.DefaultKVTableSharedCount, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *SideCarDriver) SideCarHeartBeat(heartbeat sidecartypes.SideCarHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.WebPeerID)
	var uObject, afterSetNewObj = p.SideCarTable.MustGetObject(peerID)
	var uSideCarInfo = soloboattypes.SideCarInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
	}

	uSideCarInfo.Ptr().LastHeatBeatAt = time.Now()
	uSideCarInfo.Ptr().SideCarHeartBeat = heartbeat
	err = p.FormatSideCarInfo(uSideCarInfo.Ptr())
	if err != nil {
		return err
	}

	return nil
}

func (p *SideCarDriver) FormatSideCarInfo(SideCarInfo *soloboattypes.SideCarInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(SideCarInfo.WebPeerID))
	if err != nil {
		return err
	}
	SideCarInfo.WebServerAddr = peer.AddressStr()

	return nil
}

func (p *SideCarDriver) ListObject(listPeer offheap.LKVTableListObjectWithBytes64) {
	p.SideCarTable.ListObject(listPeer)
}

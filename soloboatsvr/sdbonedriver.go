package soloboatsvr

import (
	"soloos/common/sdbapitypes"
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"soloos/soloboat/soloboattypes"
	"time"
)

type SDBOneDriver struct {
	soloBoatSvr *SoloBoatSvr
	sdbOneTable offheap.LKVTableWithBytes64
}

func (p *SDBOneDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error
	p.soloBoatSvr = soloBoatSvr
	err = p.soloBoatSvr.SoloOSEnv.OffheapDriver.InitLKVTableWithBytes64(&p.sdbOneTable, "SDBOne",
		int(soloboattypes.SDBOneInfoStructSize), -1, offheap.DefaultKVTableSharedCount, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *SDBOneDriver) SDBOneHeartBeat(heartbeat sdbapitypes.SDBOneHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.SRPCPeerID)
	var uObject, afterSetNewObj = p.sdbOneTable.MustGetObject(peerID)
	var uSDBOneInfo = soloboattypes.SDBOneInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
	}

	uSDBOneInfo.Ptr().LastHeatBeatAt = time.Now()
	uSDBOneInfo.Ptr().SDBOneHeartBeat = heartbeat
	err = p.FormatSDBOneInfo(uSDBOneInfo.Ptr())
	if err != nil {
		return err
	}

	return nil
}

func (p *SDBOneDriver) FormatSDBOneInfo(sdbOneInfo *soloboattypes.SDBOneInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	sdbOneInfo.LastHeatBeatAtStr = sdbOneInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(sdbOneInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	sdbOneInfo.SRPCServerAddr = peer.AddressStr()

	//TODO enable WebServer
	peer, _ = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(sdbOneInfo.WebPeerID))
	sdbOneInfo.WebServerAddr = peer.AddressStr()

	return nil
}

func (p *SDBOneDriver) ListObject(listPeer offheap.LKVTableListObjectWithBytes64) {
	p.sdbOneTable.ListObject(listPeer)
}

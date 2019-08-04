package soloboatsvr

import (
	"soloos/common/snettypes"
	"soloos/common/swalapitypes"
	"soloos/sdbone/offheap"
	"soloos/soloboat/soloboattypes"
	"time"
)

type SWALBrokerDriver struct {
	soloBoatSvr     *SoloBoatSvr
	swalBrokerTable offheap.LKVTableWithBytes64
}

func (p *SWALBrokerDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error
	p.soloBoatSvr = soloBoatSvr
	err = p.soloBoatSvr.SoloOSEnv.OffheapDriver.InitLKVTableWithBytes64(&p.swalBrokerTable, "SWALBroker",
		int(soloboattypes.SWALBrokerInfoStructSize), -1, offheap.DefaultKVTableSharedCount, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *SWALBrokerDriver) SWALBrokerHeartBeat(heartbeat swalapitypes.BrokerHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.SRPCPeerID)
	var uObject, afterSetNewObj = p.swalBrokerTable.MustGetObject(peerID)
	var uSWALBrokerInfo = soloboattypes.SWALBrokerInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
	}

	uSWALBrokerInfo.Ptr().LastHeatBeatAt = time.Now()
	uSWALBrokerInfo.Ptr().BrokerHeartBeat = heartbeat
	err = p.FormatSWALBrokerInfo(uSWALBrokerInfo.Ptr())
	if err != nil {
		return err
	}

	return nil
}

func (p *SWALBrokerDriver) FormatSWALBrokerInfo(swalBrokerInfo *soloboattypes.SWALBrokerInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	swalBrokerInfo.LastHeatBeatAtStr = swalBrokerInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(swalBrokerInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	swalBrokerInfo.SRPCServerAddr = peer.AddressStr()

	//TODO enable WebServer
	peer, _ = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(swalBrokerInfo.WebPeerID))
	swalBrokerInfo.WebServerAddr = peer.AddressStr()

	return nil
}

func (p *SWALBrokerDriver) ListObject(listPeer offheap.LKVTableListObjectWithBytes64) {
	p.swalBrokerTable.ListObject(listPeer)
}

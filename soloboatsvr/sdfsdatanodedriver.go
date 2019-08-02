package soloboatsvr

import (
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"soloos/soloboat/soloboattypes"
	"time"
)

type SDFSDataNodeDriver struct {
	soloBoatSvr       *SoloBoatSvr
	sdfsDataNodeTable offheap.LKVTableWithBytes64
}

func (p *SDFSDataNodeDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error
	p.soloBoatSvr = soloBoatSvr
	err = p.soloBoatSvr.SoloOSEnv.OffheapDriver.InitLKVTableWithBytes64(&p.sdfsDataNodeTable, "SDFSDataNode",
		int(soloboattypes.SDFSDataNodeInfoStructSize), -1, offheap.DefaultKVTableSharedCount, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *SDFSDataNodeDriver) SDFSDataNodeHeartBeat(heartbeat sdfsapitypes.DataNodeHeartBeat) error {
	var err error
	var peerID = snettypes.StrToPeerID(heartbeat.SRPCPeerID)
	var uObject, afterSetNewObj = p.sdfsDataNodeTable.MustGetObject(peerID)
	var uSDFSDataNodeInfo = soloboattypes.SDFSDataNodeInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
	}

	uSDFSDataNodeInfo.Ptr().LastHeatBeatAt = time.Now()
	uSDFSDataNodeInfo.Ptr().DataNodeHeartBeat = heartbeat
	err = p.FormatSDFSDataNodeInfo(uSDFSDataNodeInfo.Ptr())
	if err != nil {
		return err
	}

	return nil
}

func (p *SDFSDataNodeDriver) FormatSDFSDataNodeInfo(sdfsDataNodeInfo *soloboattypes.SDFSDataNodeInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	sdfsDataNodeInfo.LastHeatBeatAtStr = sdfsDataNodeInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(sdfsDataNodeInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	sdfsDataNodeInfo.SRPCServerAddr = peer.AddressStr()

	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(sdfsDataNodeInfo.WebPeerID))
	if err != nil {
		return err
	}
	sdfsDataNodeInfo.WebServerAddr = peer.AddressStr()

	return nil
}

func (p *SDFSDataNodeDriver) ListObject(listPeer offheap.LKVTableListObjectWithBytes64) {
	p.sdfsDataNodeTable.ListObject(listPeer)
}

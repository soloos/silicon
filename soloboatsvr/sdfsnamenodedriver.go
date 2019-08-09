package soloboatsvr

import (
	"soloos/common/snettypes"
	"soloos/sdbone/offheap"
	"soloos/soloboat/soloboattypes"
)

type SDFSNameNodeDriver struct {
	soloBoatSvr       *SoloBoatSvr
	sdfsNameNodeTable offheap.LKVTableWithBytes64
}

func (p *SDFSNameNodeDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error
	p.soloBoatSvr = soloBoatSvr
	err = p.soloBoatSvr.SoloOSEnv.OffheapDriver.InitLKVTableWithBytes64(&p.sdfsNameNodeTable, "SDFSNameNode",
		int(soloboattypes.SDFSNameNodeInfoStructSize), -1, offheap.DefaultKVTableSharedCount, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *SDFSNameNodeDriver) FormatSDFSNameNodeInfo(sdfsNameNodeInfo *soloboattypes.SDFSNameNodeInfo) error {
	var (
		peer snettypes.Peer
		err  error
	)

	sdfsNameNodeInfo.LastHeatBeatAtStr = sdfsNameNodeInfo.LastHeatBeatAt.Format("2006-01-02 15:04:05")
	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(sdfsNameNodeInfo.SRPCPeerID))
	if err != nil {
		return err
	}
	sdfsNameNodeInfo.SRPCServerAddr = peer.AddressStr()

	peer, err = p.soloBoatSvr.SNetDriver.GetPeer(snettypes.StrToPeerID(sdfsNameNodeInfo.WebPeerID))
	if err != nil {
		return err
	}
	sdfsNameNodeInfo.WebServerAddr = peer.AddressStr()

	return nil
}

func (p *SDFSNameNodeDriver) ListObject(listPeer offheap.LKVTableListObjectWithBytes64) {
	p.sdfsNameNodeTable.ListObject(listPeer)
}

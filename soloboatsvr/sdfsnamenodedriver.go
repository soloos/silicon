package soloboatsvr

import (
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SDFSNameNodeDriver struct {
	soloBoatSvr       *SoloBoatSvr
	sdfsNameNodeTable sync.Map
}

func (p *SDFSNameNodeDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	p.soloBoatSvr = soloBoatSvr
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

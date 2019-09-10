package soloboatsvr

import (
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"sync"
)

type SDFSDataNodeDriver struct {
	soloBoatSvr       *SoloBoatSvr
	sdfsDataNodeTable sync.Map
}

func (p *SDFSDataNodeDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	p.soloBoatSvr = soloBoatSvr
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

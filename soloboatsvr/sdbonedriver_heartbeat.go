package soloboatsvr

import (
	"soloos/common/sdbapitypes"
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
	"time"
)

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

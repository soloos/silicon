package soloboatsvr

import (
	"soloos/common/util"
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

func (p *SDFSNameNodeDriver) SDFSNameNodeHeartBeat(sdfsNameNode soloboattypes.SDFSNameNodeInfo) error {
	var err error
	var isNeedUpdateInDB bool = false
	var uObject, afterSetNewObj = p.sdfsNameNodeTable.MustGetObject(sdfsNameNode.ID)
	var uSDFSNameNodeInfo = soloboattypes.SDFSNameNodeInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
		// uSDFSNameNodeInfo.Ptr().SetAddress(peer.AddressStr())
		// uSDFSNameNodeInfo.Ptr().ServiceProtocol = peer.ServiceProtocol
		isNeedUpdateInDB = true
	} else {
		// isNeedUpdateInDB = uSDFSNameNodeInfo.Ptr().Address != peer.Address ||
		// uSDFSNameNodeInfo.Ptr().ServiceProtocol != peer.ServiceProtocol
	}

	util.Ignore(uSDFSNameNodeInfo)
	util.Ignore(err)
	util.Ignore(isNeedUpdateInDB)

	// if isNeedUpdateInDB {
	// err = p.StoreSDFSNameNodeInDB(*uSDFSNameNodeInfo.Ptr())
	// if err != nil {
	// return err
	// }
	// }

	return nil
}

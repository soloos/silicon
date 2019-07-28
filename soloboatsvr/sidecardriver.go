package soloboatsvr

import (
	"soloos/common/util"
	"soloos/sdbone/offheap"
	"soloos/soloboat/soloboattypes"
)

type SideCarDriver struct {
	soloBoatSvr      *SoloBoatSvr
	sideCarInfoTable offheap.LKVTableWithBytes64
}

func (p *SideCarDriver) Init(soloBoatSvr *SoloBoatSvr) error {
	var err error
	p.soloBoatSvr = soloBoatSvr
	err = p.soloBoatSvr.SoloOSEnv.OffheapDriver.InitLKVTableWithBytes64(&p.sideCarInfoTable, "SideCar",
		int(soloboattypes.SideCarInfoStructSize), -1, offheap.DefaultKVTableSharedCount, nil)
	if err != nil {
		return err
	}

	return nil
}

func (p *SideCarDriver) SidCarHeartBeat(sideCarInfo soloboattypes.SideCarInfo) error {
	var err error
	var isNeedUpdateInDB bool = false
	var uObject, afterSetNewObj = p.sideCarInfoTable.MustGetObject(sideCarInfo.ID)
	var uSideCarInfo = soloboattypes.SideCarInfoUintptr(uObject)
	if afterSetNewObj != nil {
		afterSetNewObj()
		// uSideCarInfo.Ptr().SetAddress(peer.AddressStr())
		// uSideCarInfo.Ptr().ServiceProtocol = peer.ServiceProtocol
		isNeedUpdateInDB = true
	} else {
		// isNeedUpdateInDB = uSideCarInfo.Ptr().Address != peer.Address ||
		// uSideCarInfo.Ptr().ServiceProtocol != peer.ServiceProtocol
	}

	util.Ignore(uSideCarInfo)

	if isNeedUpdateInDB {
		err = p.StoreSideCarInDB(*uSideCarInfo.Ptr())
		if err != nil {
			return err
		}
	}

	return nil
}

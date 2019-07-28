package soloboatsvr

import (
	"soloos/common/sdbapi"
	"soloos/soloboat/soloboattypes"
)

func (p *SideCarDriver) StoreSideCarInDB(sideCarInfo soloboattypes.SideCarInfo) error {
	var (
		sess sdbapi.Session
		err  error
	)

	err = p.soloBoatSvr.dbConn.InitSession(&sess)
	if err != nil {
		return err
	}

	err = sess.ReplaceInto("b_sidecar").
		PrimaryColumns("peer_id").PrimaryValues(sideCarInfo.PeerIDStr()).
		Columns("desc").
		Values(sideCarInfo.PeerIDStr()).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

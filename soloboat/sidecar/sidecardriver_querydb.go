package sidecar

import (
	"soloos/common/solodbapi"
	"soloos/soloboat/soloboattypes"
)

func (p *SidecarDriver) StoreSidecarInDB(sideCarInfo soloboattypes.SidecarInfo) error {
	var (
		sess solodbapi.Session
		err  error
	)

	err = p.soloboatIns.dbConn.InitSession(&sess)
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

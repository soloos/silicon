package soloboatsvr

import (
	"database/sql"
	"soloos/common/sdbapi"
	"soloos/common/sdfsapitypes"
	"soloos/common/snettypes"
)

type ListSNetPeerFromDB func(peer snettypes.Peer) bool

func (p *SoloBoatSvr) ListSNetPeerFromDB(listPeer ListSNetPeerFromDB) error {
	var (
		sess        sdbapi.Session
		sqlRows     *sql.Rows
		peerIDStr   string
		addrStr     string
		protocolStr string
		peer        snettypes.Peer
		err         error
	)

	err = p.dbConn.InitSession(&sess)
	if err != nil {
		goto QUERY_DONE
	}

	sqlRows, err = sess.Select("peer_id", "address", "service_protocol").
		From("b_snetpeer").
		Rows()
	if err != nil {
		goto QUERY_DONE
	}

	for sqlRows.Next() {
		err = sqlRows.Scan(&peerIDStr, &addrStr, &protocolStr)
		peer.SetPeerIDFromStr(peerIDStr)
		peer.SetAddress(addrStr)
		peer.ServiceProtocol.SetProtocolStr(protocolStr)
		if listPeer(peer) == false {
			break
		}
		if err != nil {
			goto QUERY_DONE
		}
	}

QUERY_DONE:
	if sqlRows != nil {
		sqlRows.Close()
	}
	return err
}

func (p *SoloBoatSvr) FetchSNetPeerFromDB(peerID snettypes.PeerID) (snettypes.Peer, error) {
	var (
		sess        sdbapi.Session
		sqlRows     *sql.Rows
		peer        snettypes.Peer
		addrStr     string
		protocolStr string
		err         error
	)

	err = p.dbConn.InitSession(&sess)
	if err != nil {
		goto QUERY_DONE
	}

	peer.ID = peerID
	sqlRows, err = sess.Select("address", "service_protocol").
		From("b_netinode").
		Where("peer_id=?", peer.PeerIDStr()).Rows()
	if err != nil {
		goto QUERY_DONE
	}

	if sqlRows.Next() == false {
		err = sdfsapitypes.ErrObjectNotExists
		goto QUERY_DONE
	}

	err = sqlRows.Scan(&addrStr, &protocolStr)
	if err != nil {
		goto QUERY_DONE
	}
	peer.SetAddress(addrStr)
	peer.ServiceProtocol.SetProtocolStr(protocolStr)

QUERY_DONE:
	if sqlRows != nil {
		sqlRows.Close()
	}

	return peer, err
}

func (p *SoloBoatSvr) RegisterSNetPeerInDB(peer snettypes.Peer) error {
	var (
		sess sdbapi.Session
		err  error
	)

	err = p.dbConn.InitSession(&sess)
	if err != nil {
		return err
	}

	err = sess.ReplaceInto("b_snetpeer").
		PrimaryColumns("peer_id").PrimaryValues(peer.PeerIDStr()).
		Columns("address", "service_protocol").
		Values(peer.AddressStr(), peer.ServiceProtocol.Str()).
		Exec()
	if err != nil {
		return err
	}

	return nil
}

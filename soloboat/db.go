package soloboat

import "soloos/common/solodbapi"

func (p *Soloboat) GetDBConn() *solodbapi.Connection {
	return &p.dbConn
}

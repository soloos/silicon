package soloboat

import "soloos/common/log"

func (p *Soloboat) installSchema(dbDriver string) error {
	var (
		sqls []string
		err  error
	)

	sqls = p.prepareSchemaSqls(dbDriver)
	for _, sql := range sqls {
		_, err = p.dbConn.Exec(sql)
		if err != nil {
			log.Error(err, sql)
		}
	}

	return nil
}

func (p *Soloboat) prepareSchemaSqls(dbDriver string) []string {
	var sqls []string

	sqls = append(sqls, `
	create table if not exists b_sidecar (
		peer_id char(64),
		description varchar(256),
		primary key(peer_id)
	);
	`)

	sqls = append(sqls, `
       create table if not exists b_snetpeer (
               peer_id char(64),
               address char(128),
               service_protocol char(8),
               primary key(peer_id)
       );
       `)

	// sqls = append(sqls, `
	// create index if not exists i_b_snetpeer
	// on b_snetpeer(service_protocol);
	// `)

	return sqls
}

package agent

import "soloos/common/log"

func (p *SiliconAgent) installSchema(dbDriver string) error {
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

func (p *SiliconAgent) prepareSchemaSqls(dbDriver string) []string {
	var sqls []string

	// sqls = append(sqls, `
	// create table if not exists b_silicon_conf (
	// peer_id char(64),
	// peer_type char(16),
	// conf text,
	// primary key(peer_id)
	// );
	// `)

	// sqls = append(sqls, `
	// create index if not exists i_b_silicon_conf
	// on b_silicon_conf(peer_type);
	// `)

	return sqls
}

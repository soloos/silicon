package soloboatsvr

import "soloos/common/iron"

type SoloBoatSvrOptions struct {
	SNetDriverListenAddr string
	SNetDriverServeAddr  string
	WebServerOptions     iron.Options

	PeerID   string
	DBDriver string
	Dsn      string

	PProfListenAddr string
}

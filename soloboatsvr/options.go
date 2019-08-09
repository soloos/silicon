package soloboatsvr

import "soloos/common/iron"

type SoloBoatSvrOptions struct {
	WebServerOptions iron.Options
	WebPeerID        string
	DBDriver         string
	Dsn              string

	PProfListenAddr string
}

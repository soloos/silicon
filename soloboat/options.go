package soloboat

import "soloos/common/iron"

type SoloboatOptions struct {
	WebServerOptions iron.Options
	WebPeerID        string
	DBDriver         string
	Dsn              string

	PProfListenAddr string
}

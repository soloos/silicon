package agent

import "soloos/common/iron"

type SiliconAgentOptions struct {
	SNetDriverListenAddr string
	SNetDriverServeAddr  string
	WebServerOptions     iron.Options

	PeerID   string
	DBDriver string
	Dsn      string

	PProfListenAddr string
}

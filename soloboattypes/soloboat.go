package soloboattypes

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
	"soloos/common/solodbapi"
	"soloos/common/soloosbase"
)

type Soloboat interface {
	iron.IProxy
	GetSoloosEnv() *soloosbase.SoloosEnv
	GetDBConn() *solodbapi.Connection
	GetWebServer() *iron.Server
	GetWebPeer() snettypes.Peer
}

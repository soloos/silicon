package soloboat

func (p *WebServer) prepareApi() error {
	p.server.Router("/Api/Soloboat/Sidecar/HeartBeat", p.apiSidecarHeartBeat)
	p.server.Router("/Api/SoloosCommon/Peer/List", p.apiPeerList)
	p.server.Router("/Api/Solofs/Solodn/HeartBeat", p.apiSolodnHeartBeat)
	p.server.Router("/Api/Solomq/Solomq/HeartBeat", p.apiSolomqHeartBeat)
	p.server.Router("/Api/SDB/Solodb/HeartBeat", p.apiSolodbHeartBeat)
	return nil
}

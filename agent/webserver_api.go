package agent

func (p *WebServer) prepareApi() error {
	p.server.Router("/Api/Peer/List", p.apiPeerList)
	p.server.Router("/Api/Peer/HeartBeat", p.apiPeerHeartBeat)
	return nil
}

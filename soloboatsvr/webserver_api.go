package soloboatsvr

func (p *WebServer) prepareApi() error {
	p.server.Router("/Api/Peer/List", p.apiPeerList)
	p.server.Router("/Api/SDFS/NameNode/HeartBeat", p.apiSDFSNameNodeHeartBeat)
	p.server.Router("/Api/SDFS/DataNode/HeartBeat", p.apiSDFSDataNodeHeartBeat)
	p.server.Router("/Api/SideCar/HeartBeat", p.apiSideCarHeartBeat)
	p.server.Router("/Api/Peer/HeartBeat", p.apiPeerHeartBeat)
	return nil
}

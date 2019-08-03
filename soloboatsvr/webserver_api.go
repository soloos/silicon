package soloboatsvr

func (p *WebServer) prepareApi() error {
	p.server.Router("/Api/SoloOSCommon/Peer/List", p.apiPeerList)
	p.server.Router("/Api/SDFS/NameNode/HeartBeat", p.apiSDFSNameNodeHeartBeat)
	p.server.Router("/Api/SDFS/DataNode/HeartBeat", p.apiSDFSDataNodeHeartBeat)
	p.server.Router("/Api/SWAL/Broker/HeartBeat", p.apiSWALBrokerHeartBeat)
	p.server.Router("/Api/SoloBoat/SideCar/HeartBeat", p.apiSideCarHeartBeat)
	return nil
}

package soloboatsvr

func (p *WebServer) prepareCtr() error {
	var vpre = ""
	var vcom = vpre + "/Frame.html"

	p.server.Router("/", p.ctrMain)
	p.server.AssignView("/Index", vcom, vpre+"/Index.html")

	p.server.HookBeforeHttpHandle("/SDFS/NameNode", p.prepareCtrSDFSNameNode)
	p.server.Router("/SDFS/NameNode", p.ctrSDFSNameNode)
	p.server.AssignView("/SDFS/NameNode/Index", vcom, vpre+"/SDFS/NameNode/Index.html")

	p.server.HookBeforeHttpHandle("/SDFS/DataNode", p.prepareCtrSDFSDataNode)
	p.server.Router("/SDFS/DataNode", p.ctrSDFSDataNode)
	p.server.AssignView("/SDFS/DataNode/Index", vcom, vpre+"/SDFS/DataNode/Index.html")

	p.server.HookBeforeHttpHandle("/SNet/Peer", p.prepareCtrSNetPeer)
	p.server.Router("/SNet/Peer", p.ctrSNetPeer)
	p.server.AssignView("/SNet/Peer/Index", vcom, vpre+"/SNet/Peer/Index.html")

	return nil
}

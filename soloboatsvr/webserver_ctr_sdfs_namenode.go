package soloboatsvr

import "soloos/common/iron"

func (p *WebServer) prepareCtrSDFSNameNode(ir *iron.Request) bool {
	var module = ViewModule{
		Name: "SDFSNameNode",
		CH:   "SDFSNameNode",
		URL:  "/SDFS/NameNode",
	}
	ir.ViewData["Module"] = module
	return true
}

func (p *WebServer) ctrSDFSNameNode(ir *iron.Request) {
	ir.Render("/SDFS/NameNode/Index")
}

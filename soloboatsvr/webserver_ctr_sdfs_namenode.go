package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/soloboat/soloboattypes"
)

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
	var ret []soloboattypes.SDFSNameNodeInfo
	p.soloBoatSvr.sdfsNameNodeDriver.sdfsNameNodeTable.Range(func(kIptr, vIptr interface{}) bool {
		var obj = vIptr.(soloboattypes.SDFSNameNodeInfo)
		ret = append(ret, obj)
		return true
	})
	ir.ViewData["NameNodeArr"] = ret
	ir.Render("/SDFS/NameNode/Index")
}

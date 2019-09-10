package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/soloboat/soloboattypes"
)

func (p *WebServer) prepareCtrSDFSDataNode(ir *iron.Request) bool {
	var module = ViewModule{
		Name: "SDFSDataNode",
		CH:   "SDFSDataNode",
		URL:  "/SDFS/DataNode",
	}
	ir.ViewData["Module"] = module
	return true
}

func (p *WebServer) ctrSDFSDataNode(ir *iron.Request) {
	var ret []soloboattypes.SDFSDataNodeInfo
	p.soloBoatSvr.sdfsDataNodeDriver.sdfsDataNodeTable.Range(func(kIptr, vIptr interface{}) bool {
		var obj = vIptr.(soloboattypes.SDFSDataNodeInfo)
		ret = append(ret, obj)
		return true
	})
	ir.ViewData["DataNodeArr"] = ret
	ir.Render("/SDFS/DataNode/Index")
}

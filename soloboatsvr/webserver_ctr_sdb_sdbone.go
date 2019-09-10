package soloboatsvr

import (
	"soloos/common/iron"
	"soloos/soloboat/soloboattypes"
)

func (p *WebServer) prepareCtrSDBOne(ir *iron.Request) bool {
	var module = ViewModule{
		Name: "SDBOne",
		CH:   "SDBOne",
		URL:  "/SDB/SDBOne",
	}
	ir.ViewData["Module"] = module
	return true
}

func (p *WebServer) ctrSDBOne(ir *iron.Request) {
	var ret []soloboattypes.SDBOneInfo
	p.soloBoatSvr.sdbOneDriver.sdbOneTable.Range(func(kIptr, vIptr interface{}) bool {
		var obj = vIptr.(soloboattypes.SDBOneInfo)
		ret = append(ret, obj)
		return true
	})
	ir.ViewData["SDBOneArr"] = ret
	ir.Render("/SDB/SDBOne/Index")
}

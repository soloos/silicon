package main

import (
	"soloos/common/soloosbase"
	"soloos/common/util"
	"soloos/soloboat/sidecar"
)

func main() {
	var sideCar sidecar.Sidecar
	var soloOSEnv soloosbase.SoloOSEnv
	util.AssertErrIsNil(soloOSEnv.InitWithSNet(""))
	util.AssertErrIsNil(sideCar.Init(&soloOSEnv, sidecar.Options{}))
	util.AssertErrIsNil(sideCar.Serve())
}

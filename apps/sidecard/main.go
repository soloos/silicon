package main

import (
	"soloos/common/soloosbase"
	"soloos/common/util"
	"soloos/soloboat/sidecar"
)

func main() {
	var sideCar sidecar.Sidecar
	var soloosEnv soloosbase.SoloosEnv
	util.AssertErrIsNil(soloosEnv.InitWithSNet(""))
	util.AssertErrIsNil(sideCar.Init(&soloosEnv, sidecar.Options{}))
	util.AssertErrIsNil(sideCar.Serve())
}

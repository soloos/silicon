package main

import (
	"soloos/common/util"
	"soloos/soloboat/sidecar"
)

func main() {
	var sideCar sidecar.SideCar
	util.AssertErrIsNil(sideCar.Init())
	util.AssertErrIsNil(sideCar.Serve())
}

package main

import (
	"os"
	"soloos/common/util"
	"soloos/soloboat/soloboatsvrd"
)

func main() {
	var (
		soloBoatSvrdIns soloboatsvrd.SoloBoatSvrd
		options         soloboatsvrd.Options
	)

	optionsFile := os.Args[1]
	util.AssertErrIsNil(util.LoadOptionsFile(optionsFile, &options))
	util.AssertErrIsNil(soloBoatSvrdIns.Init(options))
	util.AssertErrIsNil(soloBoatSvrdIns.Serve())
}

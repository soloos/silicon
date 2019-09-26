package main

import (
	"os"
	"soloos/common/util"
	"soloos/soloboat/soloboatd"
)

func main() {
	var (
		soloboatInsdIns soloboatd.SoloboatDaemon
		options         soloboatd.Options
	)

	optionsFile := os.Args[1]
	util.AssertErrIsNil(util.LoadOptionsFile(optionsFile, &options))
	util.AssertErrIsNil(soloboatInsdIns.Init(options))
	util.AssertErrIsNil(soloboatInsdIns.Serve())
}

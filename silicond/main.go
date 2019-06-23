package main

import (
	"os"
	"soloos/common/util"
)

func main() {
	var (
		env     Env
		options Options
		err     error
	)

	optionsFile := os.Args[1]

	err = util.LoadOptionsFile(optionsFile, &options)
	util.AssertErrIsNil(err)

	env.Init(options)
	util.AssertErrIsNil(env.Serve())
}

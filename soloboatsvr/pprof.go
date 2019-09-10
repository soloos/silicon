package soloboatsvr

import "soloos/common/util"

func (p *SoloBoatSvr) preparePProf(pprofListenAddr string) error {
	if pprofListenAddr != "" {
		go util.PProfServe(pprofListenAddr)
	}
	return nil
}

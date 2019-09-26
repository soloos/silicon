package soloboat

import "soloos/common/util"

func (p *Soloboat) preparePProf(pprofListenAddr string) error {
	if pprofListenAddr != "" {
		go util.PProfServe(pprofListenAddr)
	}
	return nil
}

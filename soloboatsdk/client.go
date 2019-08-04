package soloboatsdk

import (
	"soloos/common/snettypes"
	"soloos/common/soloosbase"
)

type Client struct {
	soloOSEnv       *soloosbase.SoloOSEnv
	soloboatWebPeer snettypes.Peer
}

func (p *Client) Init(soloOSEnv *soloosbase.SoloOSEnv,
	soloboatWebPeerID snettypes.PeerID) error {
	var err error

	p.soloOSEnv = soloOSEnv
	p.soloboatWebPeer, err = p.soloOSEnv.SNetDriver.GetPeer(soloboatWebPeerID)
	if err != nil {
		return err
	}
	return nil
}

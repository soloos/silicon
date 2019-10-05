package soloboatsdk

import (
	"soloos/common/snettypes"
	"soloos/common/soloosbase"
)

type Client struct {
	soloosEnv       *soloosbase.SoloosEnv
	soloboatWebPeer snettypes.Peer
}

func (p *Client) Init(soloosEnv *soloosbase.SoloosEnv,
	soloboatWebPeerID snettypes.PeerID) error {
	var err error

	p.soloosEnv = soloosEnv
	p.soloboatWebPeer, err = p.soloosEnv.SNetDriver.GetPeer(soloboatWebPeerID)
	if err != nil {
		return err
	}
	return nil
}

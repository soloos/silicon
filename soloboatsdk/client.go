package soloboatsdk

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
)

type Client struct {
	soloboatServeAddr string
}

func (p *Client) Init(soloboatServeAddr string) error {
	p.soloboatServeAddr = soloboatServeAddr
	return nil
}

func (p *Client) HeartBeat(peerID snettypes.PeerID) error {
	var (
		urlPath = p.soloboatServeAddr + "/Api/Peer/HeartBeat"
		req     HeartBeatReqJSON
		resp    HeartBeatRespJSON
	)
	req.PeerID = peerID.Str()
	return iron.PostJSON(urlPath, req, &resp)
}

package siliconsdk

import (
	"soloos/common/iron"
	"soloos/common/snettypes"
)

type Client struct {
	siliconServeAddr string
}

func (p *Client) Init(siliconServeAddr string) error {
	p.siliconServeAddr = siliconServeAddr
	return nil
}

func (p *Client) HeartBeat(peerID snettypes.PeerID) error {
	var (
		urlPath = p.siliconServeAddr + "/Peer/HeartBeat"
		req     HeartBeatReqJSON
		resp    HeartBeatRespJSON
	)
	req.PeerID = peerID.Str()
	return iron.PostJSON(urlPath, req, &resp)
}

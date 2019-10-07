package soloboatsdk

import "soloos/common/snet"

type HeartBeatReq struct {
	PeerID string
}

type HeartBeatResp struct {
	snet.RespDataCommon
}

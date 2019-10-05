package soloboatsdk

import "soloos/common/snettypes"

type HeartBeatReq struct {
	PeerID string
}

type HeartBeatResp struct {
	snettypes.RespDataCommon
}

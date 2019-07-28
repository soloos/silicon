package soloboatsdk

import "soloos/common/snettypes"

type HeartBeatReqJSON struct {
	PeerID string
}

type HeartBeatRespJSON struct {
	snettypes.APIRespCommonJSON
}

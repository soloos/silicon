package soloboatsdk

import "soloos/common/snettypes"

type SDFSNameNodeHeartBeatReqJSON struct {
	PeerID string
}

type SDFSNameNodeHeartBeatRespJSON struct {
	snettypes.APIRespCommonJSON
}

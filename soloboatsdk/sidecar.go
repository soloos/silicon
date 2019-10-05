package soloboatsdk

import (
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
)

type SidecarInfoReq struct {
	soloboattypes.SidecarInfoJSON
}

type SidecarInfoResp struct {
	snettypes.RespDataCommon
}

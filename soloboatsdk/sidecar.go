package soloboatsdk

import (
	"soloos/common/snet"
	"soloos/soloboat/soloboattypes"
)

type SidecarInfoReq struct {
	soloboattypes.SidecarInfoJSON
}

type SidecarInfoResp struct {
	snet.RespDataCommon
}

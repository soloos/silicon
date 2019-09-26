package soloboatsdk

import (
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
)

type SidecarInfoReqJSON struct {
	soloboattypes.SidecarInfoJSON
}

type SidecarInfoRespJSON struct {
	snettypes.APIRespCommonJSON
}

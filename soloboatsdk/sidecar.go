package soloboatsdk

import (
	"soloos/common/snettypes"
	"soloos/soloboat/soloboattypes"
)

type SideCarInfoReqJSON struct {
	soloboattypes.SideCarInfoJSON
}

type SideCarInfoRespJSON struct {
	snettypes.APIRespCommonJSON
}

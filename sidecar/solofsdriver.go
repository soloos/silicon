package sidecar

import "soloos/common/solofsapitypes"

type SolofsDriver struct {
	sideCar *Sidecar
}

func (p *SolofsDriver) Init(sideCar *Sidecar) error {
	p.sideCar = sideCar
	return nil
}

func (p *SolofsDriver) TraceNetBlock(netINodeID solofsapitypes.NetINodeID, netBlockIndex int) error {
	return nil
}

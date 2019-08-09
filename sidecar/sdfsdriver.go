package sidecar

import "soloos/common/sdfsapitypes"

type SDFSDriver struct {
	sideCar *SideCar
}

func (p *SDFSDriver) Init(sideCar *SideCar) error {
	p.sideCar = sideCar
	return nil
}

func (p *SDFSDriver) TraceNetBlock(netINodeID sdfsapitypes.NetINodeID, netBlockIndex int) error {
	return nil
}

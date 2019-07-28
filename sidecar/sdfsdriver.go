package sidecar

import "soloos/common/sdfsapitypes"

type SDFSDriver struct {
	sideCar *SideCar
	options Options
}

func (p *SDFSDriver) Init(sideCar *SideCar, options Options) error {
	p.sideCar = sideCar
	p.options = options
	return nil
}

func (p *SDFSDriver) TraceNetBlock(netINodeID sdfsapitypes.NetINodeID, netBlockIndex int) error {
	return nil
}

package sidecar

type SWALDriver struct {
	sideCar *SideCar
}

func (p *SWALDriver) Init(sideCar *SideCar) error {
	p.sideCar = sideCar
	return nil
}

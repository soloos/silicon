package sidecar

type SWALDriver struct {
	sideCar *SideCar
	options Options
}

func (p *SWALDriver) Init(sideCar *SideCar, options Options) error {
	p.sideCar = sideCar
	p.options = options
	return nil
}

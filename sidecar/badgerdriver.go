package sidecar

type BadgerDriver struct {
	sideCar *SideCar
	options Options
}

func (p *BadgerDriver) Init(sideCar *SideCar, options Options) error {
	p.sideCar = sideCar
	p.options = options
	return nil
}

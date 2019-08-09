package sidecar

type BadgerDriver struct {
	sideCar *SideCar
}

func (p *BadgerDriver) Init(sideCar *SideCar) error {
	p.sideCar = sideCar
	return nil
}

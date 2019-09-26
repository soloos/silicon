package sidecar

type BadgerDriver struct {
	sideCar *Sidecar
}

func (p *BadgerDriver) Init(sideCar *Sidecar) error {
	p.sideCar = sideCar
	return nil
}

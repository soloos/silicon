package sidecar

type SolomqDriver struct {
	sideCar *Sidecar
}

func (p *SolomqDriver) Init(sideCar *Sidecar) error {
	p.sideCar = sideCar
	return nil
}

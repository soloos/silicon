package solonn

func (p *SolonnDriver) initService() error {
	p.soloboat.RegisterService("/Solonn/Heatbeat", p.SolonnHeartBeat)
	return nil
}

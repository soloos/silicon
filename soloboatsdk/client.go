package soloboatsdk

type Client struct {
	soloboatServeAddr string
}

func (p *Client) Init(soloboatServeAddr string) error {
	p.soloboatServeAddr = soloboatServeAddr
	return nil
}

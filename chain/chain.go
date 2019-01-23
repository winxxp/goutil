package chain

type Handle func() error

func Run(h ...Handle) error {
	return New().Handles(h...).Run()
}

type HandleChain struct {
	err     error
	handles []Handle
}

func New() *HandleChain {
	return &HandleChain{
		handles: make([]Handle, 0, 3),
	}
}

func (r *HandleChain) Handles(h ...Handle) *HandleChain {
	r.handles = append(r.handles, h...)
	return r
}

func (c *HandleChain) Run() error {
	for _, h := range c.handles {
		if c.err == nil {
			c.err = h()
		}
	}

	return c.err
}

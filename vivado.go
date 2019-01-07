package vivado

type Vivado struct {
	location string
}

func NewVivado(location string) *Vivado {
	return &Vivado{
		location: location,
	}
}

func (v *Vivado) Batch(location string) error {
	return execBatch(v.location, location)
}

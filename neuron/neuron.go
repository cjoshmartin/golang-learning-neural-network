package neuron

type Neuron struct {
}

type Connection struct {
	weight      float64
	deltaWeight float64
}

func New() Neuron {
	// TODO

	return Neuron{}
}

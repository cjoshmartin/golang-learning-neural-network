package neuron

type Neuron struct {
	eta          float64
	alpha        float64
	outputVal    float64
	outputWeight []float64
	myIndex      uint
	gradient     float64
}

type Connection struct {
	weight      float64
	deltaWeight float64
}

func New() Neuron {
	// TODO

	return Neuron{}
}

// func randomWeight() float64{
// return rand()/float64(RA) // TODO
// }

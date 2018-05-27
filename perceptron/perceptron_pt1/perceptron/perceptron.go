package perceptron

import "math/rand"

type Perceptron struct {
	weigths []float64
}

func New(size int) Perceptron {
	var newPerceptron Perceptron
	newPerceptron.weigths = make([]float64, size)

	for i := 0; i < len(newPerceptron.weigths); i++ {
		newPerceptron.weigths[i] = (rand.Float64() * 2) - 1 // trying to send a random number [-1,1]
	}

	return newPerceptron
}

// creates a Guess
func (currentPerceptron *Perceptron) Guess(inputs []float64) int64 {
	sum := 0.0
	for i := 0; i < len(inputs); i++ {
		sum += inputs[i] * currentPerceptron.weigths[i]
	}
	return signum(sum) // passing sum to activation function
}

func signum(n float64) int64 { // activation function
	if n >= 0 {
		return 1
	}
	return -1
}

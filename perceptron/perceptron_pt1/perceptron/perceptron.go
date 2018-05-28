package perceptron

import "math/rand"

type TrainingSet struct {
	X      float64
	Y      float64
	Bias   float64
	Answer int64
}

type Perceptron struct {
	weigths      []float64
	learningRate float64
	GlobalError  float64
}

func New(size int, _learningRate float64) Perceptron {
	var newPerceptron Perceptron
	newPerceptron.weigths = make([]float64, size)

	for i := 0; i < len(newPerceptron.weigths); i++ {
		newPerceptron.weigths[i] = (rand.Float64() * 2) - 1 // trying to send a random number [-1,1]
	}
	newPerceptron.learningRate = _learningRate /* setting the learnig rate, to what has been pasted */
	return newPerceptron
}

//Guess will creates a Guess || feedForward algorthim
func (currentPerceptron *Perceptron) Guess(set TrainingSet) int64 {
	sum := 0.0
	inputs := []float64{set.X, set.Y, set.Bias}

	for i := 0; i < len(inputs); i++ {
		sum += inputs[i] * currentPerceptron.weigths[i]
	}
	return signum(sum) // passing sum to activation function
}

func (currentPerceptron *Perceptron) Train(set TrainingSet, expected int64) {

	algorthimGuess := currentPerceptron.Guess(set)
	currentPerceptron.GlobalError = float64(expected - algorthimGuess)
	inputs := []float64{set.X, set.Y, set.Bias}

	for i := range currentPerceptron.weigths {
		currentPerceptron.weigths[i] += currentPerceptron.learningRate * currentPerceptron.GlobalError * inputs[i]
	}

}

/* Utls */
func signum(n float64) int64 { // activation function
	if n >= 0 {
		return 1
	}
	return -1
}

func (currentPerceptron *Perceptron) GetWeight() []float64 {
	return currentPerceptron.weigths
}

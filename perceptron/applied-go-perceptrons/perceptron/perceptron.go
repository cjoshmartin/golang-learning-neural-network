package perceptron

import (
	"math/rand"

	"math"
)

type Perceptron struct {
	weights []float32
	bais    float32
}

/*
	Create a new perceptron with n inputs. Weights and bias
	are initialized with random values between -1 and 1.
*/

func New(size int32) Perceptron {
	var newPerceptron Perceptron
	newPerceptron.weights = make([]float32, size, size)
	newPerceptron.bais = rand.Float32() * 2 - 1

	for i := 0; int32(i) < size; i++ {
		newPerceptron.weights[i] = rand.Float32() * 2 - 1
	}

	return newPerceptron
}

/*
	Process implements the core functionality of the perceptron.
	It weighs the input signals, sums them up, adds the bias,
	and runs the result through the Heaviside Step function.
	(The return value could be a boolean but is an int32 instead,
	so that we can directly use the value for adjusting the perceptron.)
*/

func (currentPerceptron *Perceptron) FeedForward(inputs []int32) float64 { // Guess
	sum := currentPerceptron.bais

	for i, input := range inputs {
		sum += float32(input) * currentPerceptron.weights[i]
	}

	return currentPerceptron.signumActivationFunc(sum)
}

/*
	During the learning phase, the perceptron adjusts the weights and the
	bias based on how much the perceptron’s answer differs from the correct answer.
*/
func (currentPerceptron *Perceptron) Backpropagation(inputs []int32, delta float64, learningRate float32) {

	for i, input:= range inputs {
		currentPerceptron.weights[i] += float32(input) * float32(delta) * learningRate
	}

	currentPerceptron.bais += float32(delta) * learningRate
}

func (p *Perceptron) heavisideActivationFunc(x float32) int32 {
	if x < 0 {
		return 0
	}
	return 1
}

func (p *Perceptron) signumActivationFunc(x float32) float64{
	return 1.0 / (1.0 + math.Exp(-float64(x)))
}
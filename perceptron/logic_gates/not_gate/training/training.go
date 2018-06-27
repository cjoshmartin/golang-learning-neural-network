package training

import (
	"../perceptron"
	"math/rand"
)

type Training struct {

	currentPerceptron *perceptron.Perceptron
}

func New(currentPerceptron *perceptron.Perceptron) Training{

	return Training{
		currentPerceptron,
	}
}

func negateValue(point int32) int32 {

	if point == 1 {
		return 0
	}

	return 1
}
/*
Function train is our teacher. The teacher generates
 random test points and feeds them to the perceptron.
  Then the teacher compares the answer against 
  the solution from the ‘solution manual’ and tells 
  the perceptron how far it is off.
*/
func(currentTraining *Training) Train( iters int, rate float32) {

	for i := 0; i < iters; i++ {
		point := rand.Int31n(2) - 1

		actual := currentTraining.currentPerceptron.FeedForward(point)
		expected := float64(point)
		delta := expected - actual
		currentTraining.currentPerceptron.Backpropagation(point,delta, rate)
	}
}

func(currentTraining *Training) Verify() int32 {
	correctAnswer := int32(0)
	for i := 0; i < 100; i++ {

		point := rand.Int31n(2) - 1

		result := currentTraining.currentPerceptron.FeedForward(point)

		if int32(result) == negateValue(point) {
			correctAnswer += 1
		}
	}



	return correctAnswer
}

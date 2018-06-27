package training

import (
	"../perceptron"
	"math/rand"
)

type Training struct {
	a int32
	b int32
	currentPerceptron *perceptron.Perceptron
}

func New(currentPerceptron *perceptron.Perceptron) Training{


	return Training{
		rand.Int31n(11) - 6,
		rand.Int31n(101)- 51,
		currentPerceptron,
	}
}
/*
This function describes the separation line.
*/
func(currentTraining *Training) F(x int32) int32 {
	return currentTraining.a * x + currentTraining.b
}

/*
Function isAboveLine returns 1 if the point (x,y)
 is above the line y = ax + b, else 0. This is
 our teacher’s solution manual.
*/
func IsAboveLine(point []int32, f func(int32) int32) float64 {
	//x := point[0]
	//y := point[1]
	//fx := f(x)
	//
	////if y > fx {
	//	return 1
	//}
	return 0
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
		point := []int32{
			rand.Int31n(2) - 1,
		}
		actual := currentTraining.currentPerceptron.FeedForward(point)
		expected := float64(point[0])
		delta := expected - actual
		currentTraining.currentPerceptron.Backpropagation(point,delta, rate)
	}
}

func(currentTraining *Training) Verify() int32 {
	correctAnswer := int32(0)

	//canvas := draw.NewCanvas()

	for i := 0; i < 100; i++ {
		point := []int32{
			rand.Int31n(2) - 1,
		}

		result := currentTraining.currentPerceptron.FeedForward(point)

		if result == IsAboveLine(point,currentTraining.F) {
			correctAnswer += 1
		}

		//canvas.DrawPoint(point[0],point[1],result==1)
	}



	return correctAnswer
}

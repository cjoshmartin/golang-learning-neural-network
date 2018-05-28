package main

import (
	"fmt"
	"math/rand"

	"./perceptron"
)

type settings struct {
	xmin  float64
	ymin  float64
	xmax  float64
	ymax  float64
	count int64
}
type Line struct {
	X1 float64
	y1 float64
	X2 float64
	y2 float64
}

func f(x float64) float64 {
	return 0.4*x + 1
}
func NewLine(x1, y1, x2, y2 float64) Line {
	return Line{x1, y1, x2, y2}
}

func (currentLine *Line) UpdateLine(weights []float64) {
	// Draw the line based on the current weights
	// Formula is weights[0]*x + weights[1]*y + weights[2] = 0
	*currentLine = NewLine(currentLine.X1, (-weights[2]-weights[0]*currentLine.X1)/weights[1], currentLine.X2, (-weights[2]-weights[0]*currentLine.X2)/weights[1])

}

func main() {
	networkSettings := settings{-400.0, -100.0, 400.0, 100.0, 0}
	trainer := make([]perceptron.TrainingSet, 2000)

	// @parms: 1) 3 inputs -- x,y, and bias
	// @parms: 2) "Learning Constant"
	ptron := perceptron.New(3, 0.00001)

	for i := range trainer { // creating random training points and calculating the "known"
		x := (rand.Float64() * networkSettings.xmax) + networkSettings.xmin
		y := (rand.Float64() * networkSettings.ymax) + networkSettings.ymin
		answer := int64(1)
		if y < f(x) {
			answer = -1
		}
		trainer[i] = perceptron.TrainingSet{x, y, 1, answer}
	}

	lineOfApproximate := NewLine(networkSettings.xmin, f(networkSettings.xmin), networkSettings.xmax, f(networkSettings.xmax))
	for true {
		weights := ptron.GetWeight()

		lineOfApproximate.UpdateLine(weights)
		ptron.Train(trainer[networkSettings.count], trainer[networkSettings.count].Answer)
		networkSettings.count = (networkSettings.count + 1) % int64(len(trainer))

		for i := 0; i < int(networkSettings.count); i++ {
			/*guess := */ ptron.Guess(trainer[i])
			fmt.Println(ptron.GlobalError)
		}
	}
}

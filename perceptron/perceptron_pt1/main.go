package main

import (
	"fmt"

	"./perceptron"
	"./training"
)

func main() {
	inputs := []float64{-1.0, 0.5}
	test := perceptron.New(2)

	fmt.Println(test.Guess(inputs))

	trainings := training.New(100)
	trainings.Show()
}

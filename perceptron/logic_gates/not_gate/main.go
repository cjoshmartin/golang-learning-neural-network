package main

import (
	"./perceptron"
	"./training"
	"fmt"
	"time"
	"math/rand"
)
func main() {

	rand.Seed(time.Now().UnixNano())

	percept := perceptron.New(1)
	trainer := training.New(&percept)
	iterations := 1000
	learningRate := float32(0.1)

		trainer.Train(iterations, learningRate)
		successRate := trainer.Verify()

		fmt.Printf("%d%% of the answers were correct.\n", successRate)
}

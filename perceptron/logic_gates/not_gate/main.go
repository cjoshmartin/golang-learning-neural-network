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

	percept := perceptron.New()
	trainer := training.New(&percept)
	iterations := 1000
	learningRate := float32(0.1)

	for successRate := trainer.Verify();
		successRate < 80;
	    successRate = trainer.Verify() {
		fmt.Printf("%d%% of the answers were correct.\n", successRate)
		trainer.Train(iterations, learningRate)
	}
}

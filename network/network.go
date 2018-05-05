package network

import "math/rand"

type Neuron struct {
	outputValue   float64
	outputWeights []float64
	outputDeltas  []float64
}
type NeuronSettings struct {
	NumInputNeurons  uint64
	NumHiddenNeurons uint64
	NumOutputNeurons uint64
}
type InitalConditons struct {
	LearningRate float64
	Momentum     float64
	MaxEpochs    float64
	MaxError     float64
	NeuronOpt    NeuronSettings
}

type Network struct {
	inputLayer   []Neuron // where the inputs of our network go
	hiddenLayer  []Neuron // The hidden layers' job is to transform the inputs into something that the output layer can use.
	outputLayer  []Neuron // outs of our network
	trainingData []uint64 // data to train over
}

func inptLayer(settings NeuronSettings) []Neuron {
	var inputLayer []Neuron

	inputLayer = make([]Neuron, settings.NumInputNeurons+1)
	for i := range inputLayer {

		inputLayer[i].outputWeights = make([]float64, settings.NumHiddenNeurons+1)

		for j := range inputLayer[i].outputWeights {
			inputLayer[i].outputWeights[j] = rand.Float64()*(0.01-0.001) + 0.001
		} // end of outputWeights for loop

		inputLayer[i].outputDeltas = make([]float64, settings.NumHiddenNeurons)

		// setting bias neuron's output
		if i == 0 {
			inputLayer[i].outputValue = 1.0
		} else {
			//else setting output to 0
			inputLayer[i].outputValue = 0.0
		} // end of if's for setting bais

	} // end of inputLayer init foor loop
	return inputLayer
}

func New(setter InitalConditons, data []uint64) Network {
	var newNetwork Network
	// making the layers
	newNetwork.inputLayer = inptLayer(setter.NeuronOpt)
	newNetwork.hiddenLayer = make([]Neuron, setter.NeuronOpt.NumHiddenNeurons+1)
	newNetwork.outputLayer = make([]Neuron, setter.NeuronOpt.NumOutputNeurons+1)

	// sending the data as trainingData
	newNetwork.trainingData = data
	return newNetwork
}

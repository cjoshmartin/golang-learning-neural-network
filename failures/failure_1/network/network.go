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
	MaxEpochs    int
	MaxError     float64
	NeuronOpt    NeuronSettings
}

type Network struct {
	inputLayer   []Neuron // where the inputs of our network go
	hiddenLayer  []Neuron // The hidden layers' job is to transform the inputs into something that the output layer can use.
	outputLayer  []Neuron // outs of our network
	trainingData []uint64 // data to train over
}

func initIndices(length int) []int {
	indices := make([]int, length)
	for i := range indices {
		indices[i] = i
	}
	return indices
}

func layerCreate(size uint64) []Neuron {
	var Layer []Neuron

	Layer = make([]Neuron, size+1)
	for i := range Layer {

		Layer[i].outputWeights = make([]float64, size+1)

		for j := range Layer[i].outputWeights {
			Layer[i].outputWeights[j] = rand.Float64()*(0.01-0.001) + 0.001
		} // end of outputWeights for loop

		Layer[i].outputDeltas = make([]float64, size)

		// setting bias neuron's output
		if i == 0 {
			Layer[i].outputValue = 1.0
		} else {
			//else setting output to 0
			Layer[i].outputValue = 0.0
		} // end of if's for setting bais

	} // end of inputLayer init foor loop
	return Layer
}
func (network *Network) globalError(data []uint64, settings InitalConditons) float64 {
	var totalError float64

	for i := range data {
		trainData := data[i][:settings.NeuronOpt.NumInputNeurons] // Change to accomodate different data-length
		targetData := data[i][settings.NeuronOpt.NumInputNeurons:]
		network.feedForward(trainData) //todo

		for j := range network.outputLayer {
			diff := targetData[j] - network.outputLayer[j].outputValue
			totalError += diff * diff
		}
	}

	return totalError / float64((len(network.outputLayer) * len(data)))
}

func New(setter InitalConditons, data []uint64) Network {
	var newNetwork Network
	// making the layers
	newNetwork.inputLayer = layerCreate(setter.NeuronOpt.NumHiddenNeurons)
	newNetwork.hiddenLayer = layerCreate(setter.NeuronOpt.NumOutputNeurons)
	newNetwork.outputLayer = make([]Neuron, setter.NeuronOpt.NumOutputNeurons+1)

	// Output layer has no outgoing weights
	for i := range newNetwork.outputLayer {
		newNetwork.outputLayer[i].outputValue = 0
		newNetwork.outputLayer[i].outputWeights = nil
		newNetwork.outputLayer[i].outputDeltas = nil
	}
	// sending the data as trainingData
	newNetwork.trainingData = data
	return newNetwork
}

func (network *Network) Train(settings InitalConditons) int {

	indices := initIndices(len(network.trainingData))

	var epoch int
	for epoch = 0; epoch < settings.MaxEpochs; epoch++ {
		if globalError(network.trainingData, settings) < settings.MaxError {
			break
		}
	}

	return epoch
}

func (network *Network) feedForward(inputs []float64) {
	//TODO
}

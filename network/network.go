package network

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

func New(setter InitalConditons, data []uint64) Network {
	var newNetwork Network
	// making the layers
	newNetwork.inputLayer = make([]Neuron, setter.NeuronOpt.NumInputNeurons+1)
	newNetwork.hiddenLayer = make([]Neuron, setter.NeuronOpt.NumHiddenNeurons+1)
	newNetwork.outputLayer = make([]Neuron, setter.NeuronOpt.NumOutputNeurons+1)

	// sending the data as trainingData
	newNetwork.trainingData = data
	return newNetwork
}

package network
import (
	"../layer"
)

type Network struct {
	layers []layer.Layer
	error   float64
	recentAverageError float64
recentAverageSmoothingFactor float64
}

func New(topology []uint64) Network{
	var _layers []layer.Layer
	_numLayers := len(topology)

	for layerNum := 0; layerNum < _numLayers; layerNum++{
		_layers = append(_layers, layer.New())

		numOutputs := topology[layerNum +1]
		if layerNum == (len(topology) -1 ){
			numOutputs = 0
		}

		currLayer:=_layers[len(_layers)-1]
		currLayer[len(currLayer)-1].setOutputVal(1.0);
	}
	return Network{_layers,100.0}
}
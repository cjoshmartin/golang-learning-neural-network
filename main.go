package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"

	"./trainingdata"
	"./network"
)
// makes random test values for test the network
func showVectorVals(label string, v []float64) {

	fmt.Printf("%s ", label)
	for i := 0; i < len(v); i++ {
		fmt.Printf("%f ", v[i])
	}
	fmt.Printf("\n")
}
// makes a text file to test with
func genrateData(storepath string) {
	var buffer bytes.Buffer

	buffer.WriteString("topology: 2 4 1\n")

	for i := 2000; i >= 0; i-- {
		n1 := rand.Intn(2)
		n2 := rand.Intn(2)
		t := n1 ^ n2

		buffer.WriteString(fmt.Sprintf("in: %d.0 %d.0 \nout %d\n", n1, n2, t))
	}

	outputdata := []byte(buffer.String())

	ioutil.WriteFile(storepath,
		outputdata,
		0644)
}

func main() {
	filename := "sample_data.txt"
	var topology []uint64
	genrateData(filename)

	datarunner := trainingdata.New(filename)
	datarunner.OpenFile()
	datarunner.GetTopology(&topology)

	neuralnetwork := network.New(topology)
	fmt.Println(neuralnetwork)
}

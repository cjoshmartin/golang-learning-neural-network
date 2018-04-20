package trainingdata

import (
	"fmt"
	"io/ioutil"
)

type TrainingData struct {
	FileName string
	datafile []byte
}

func New(fileName string) TrainingData {
	//constructor for this "class"

	t := TrainingData{fileName, nil}

	return t
}

func (t TrainingData) OpenFile() {

	if len(t.FileName) < 0 {
		panic("Filename are not set to anything")
	}
	datafile, err := ioutil.ReadFile(t.FileName)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("All good")
	}

	t.datafile = datafile
}

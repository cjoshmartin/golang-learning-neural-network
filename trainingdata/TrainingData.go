package trainingdata

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type TrainingData struct {
	FileName string
	datafile *os.File
}

func New(fileName string) TrainingData {
	//constructor for this "class"

	t := TrainingData{fileName, nil}

	return t
}

func (t *TrainingData) OpenFile() {

	if len(t.FileName) < 0 {
		panic("Filename are not set to anything")
	}
	_datafile, err := os.Open(t.FileName)
	if err != nil {
		panic(err)
	}
	// defer _datafile.Close()
	t.datafile = _datafile
}

// GetTopology has cool stuff
func (t *TrainingData) GetTopology(topology []uint64) {

	// var _line string
	// var _label string

	scanner := bufio.NewScanner(t.datafile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

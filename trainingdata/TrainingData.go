package trainingdata

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"strconv"
	"os"
)

type TrainingData struct {
	FileName string
	datafile *os.File
	isTopologal bool
}

func New(fileName string) TrainingData {
	//constructor for this "class"
	return TrainingData{fileName, nil, false}
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
func (t *TrainingData) GetTopology(topology *[]uint64) {
	scanner := bufio.NewScanner(t.datafile)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if 	!scanner.Scan() && !strings.Contains(scanner.Text(), "topology:"){
		panic("`topology:` not found in text file")
	}else {
		t.isTopologal = true
	}

	_line := strings.Split(scanner.Text()," ")

	for i:= 1; i< len(_line); i++ {
		topologyNums, err := strconv.ParseInt(_line[i],10,64)
    if err != nil {
		panic(err)
    }
		*topology = append(*topology,uint64(topologyNums))

	} //end of for loop

	fmt.Println(topology)
} // end of getTopology

	// if 	!scanner.Scan() && !strings.Contains(scanner.Text(), "topology:"){
		
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
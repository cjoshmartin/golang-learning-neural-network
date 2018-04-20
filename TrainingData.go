package trainingData

import (
    "io/ioutil"
)

type TrainingData struct {

    FileName    string
    data_file        []byte

}

func New(fileName string) TrainingData {
    //constructor for this "class" 

    t := TrainingData{fileName,nil}

    return t
}

func(t TrainingData) OpenFile(){

    if len(t.Filename) < 0 {
        panic("Filename are not set to anything")
    }
    data_file, err := ioutil.ReadFile(t.Filename)
    if err != nil {
        panic(err)
    }

    t.data_file = data_file
}

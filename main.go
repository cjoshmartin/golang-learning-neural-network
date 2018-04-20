package main


import (
    "fmt"
    "math/rand"
    "bytes"
    "io/ioutil"
)

func showVectorVals(label string, v []float64 ){

    fmt.Printf("%s ",label)
    for  i := 0; i< len(v); i++{
        fmt.Printf("%f ",v[i])
    }
    fmt.Printf("\n")
}


func genrateData(store_path string) {
    var buffer bytes.Buffer


    buffer.WriteString("topology: 2 4 1\n")

    for i:=2000; i>=0;i--{
        n1 :=rand.Intn(2)
        n2 :=rand.Intn(2)
        t := n1 ^ n2;

        buffer.WriteString(fmt.Sprintf("in: %d.0 %d.0 \nout %d\n",n1,n2,t))
    }

    output_data :=[]byte(buffer.String())

    ioutil.WriteFile(store_path,
    output_data,
    0644)
}

func main() {
    file_name :="sample_data.txt"

    test:= "Hello tester, we be testing all day\n"
    fmt.Println(test)
    genrateData(file_name)

}

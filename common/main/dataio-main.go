package main

import (
    "argos/common/dataio"
    "fmt"
 
    "math"
    "encoding/binary"
) 
func main() {
    out := dataio.NewDataOutputX()
    out.WriteInt(10)
    var value int
    out.ReadInt(&value)
    
    fmt.Println("the value of  a : %d" , value)
    

 //   buf := new(bytes.Buffer)
    
	var pi float64 = math.Pi
	//err := binary.Write(buf, binary.LittleEndian, pi)
    err := binary.Write(out.Buffer1, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", out.Buffer1.Bytes())
}
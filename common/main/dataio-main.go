package main

import (
    "argos/common/dataio"
    "fmt"
) 
func main() {
    out := dataio.NewDataOutputX()
    out.WriteInt32(10)
    var value int32
    out.ReadInt32(&value)
    
    fmt.Println("the value of  a : %d" , value)
    


}
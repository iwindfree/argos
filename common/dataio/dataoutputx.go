package dataio

import (
    "encoding/binary"
    "bytes"
    "fmt"

)

/*var INT3_MIN_VALUE int = 0xff800000
var INT3_MAX_VALUE int = 0x007fffff
var LONG5_MIN_VALUE int64 = 0xffffff8000000000
var	LONG5_MAX_VALUE int64 = 0x0000007fffffffff */

type dataOutputX struct {
    written int
    buffer *bytes.Buffer
    Buffer1 *bytes.Buffer
}


func NewDataOutputX() *dataOutputX {
    out := new (dataOutputX)
    out.written = 0;
    out.buffer = new (bytes.Buffer)
    out.Buffer1 = new (bytes.Buffer)
    return out;
}



func (out *dataOutputX) WriteInt32(value int32) *dataOutputX {
    out.written += 4;
    err  := binary.Write(out.buffer,binary.BigEndian, value)
    if err != nil {
        fmt.Println("Failed to binary write : " ,err)
    }
    return out;
    
}

func (out *dataOutputX) ReadInt32(value *int32) {
    err:= binary.Read(out.buffer,binary.BigEndian, value)
    if err != nil {
        fmt.Println("Faileed to binary read :", err)
        value = nil
    }
    
}



package dataio
import (
   "testing"
  
)

func TestWriteInt(t *testing.T) {
    out := NewDataOutputX()
    out.WriteInt(100)
    var a int
    out.ReadInt(&a)
    t.Log("the value of a : %d", a)
}


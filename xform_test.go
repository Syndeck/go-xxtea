package xxtea

import (
	"bytes"
	"testing"
)

func TestTransform(t *testing.T) {
	b := [...]byte{0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 3, 0, 0, 0, 4, 0, 0, 0, 5}
	u := [...]int32{1, 2, 3, 4, 5}

	if g := int32ToBytes(u[:]); bytes.Compare(g, b[:]) != 0 {
		t.Errorf("convertion []int -> []byte failed:: %+v", g)
	}

	if g := bytesToInt32(b[:]); len(g) != len(u) {
		t.Errorf("convertion []byte -> []uint failed:: %+v", g)
	} else {
		for i := range g {
			if g[i] != u[i] {
				t.Errorf("convertion []byte -> []uint failed:: %+v", g)
				break
			}
		}
	}
}

package protoparse

import "testing"

func TestUnpackXiaMen(t *testing.T) {
	UnpackXiaMen([]byte{71, 71, 21, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 145, 2, 20, 0, 0, 0, 0, 0, 0, 0, 94, 19}, nil, nil, nil)
}

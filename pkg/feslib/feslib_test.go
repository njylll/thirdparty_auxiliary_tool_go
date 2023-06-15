package feslib

import (
	"constdefine"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type calcTest struct {
	a uint32
	b []byte
}

var byteTests = []calcTest{
	calcTest{131328, []byte{0, 1, 2, 0}},
	calcTest{513, []byte{1, 2, 0, 0}},
}

func TestByteToUint(t *testing.T) {

	for _, v := range byteTests {
		ret := ByteToUint(v.b)
		if ret != v.a {
			t.Errorf("%v  want %d,but get %d", v.b, v.a, ret)
		}
	}

}

func TestUintToByte(t *testing.T) {
	for _, v := range byteTests {
		ret := UintToByte(v.a)
		if !reflect.DeepEqual(ret, v.b) {
			t.Errorf("%v  want %d,but get %d", v.a, v.b, ret)
		}
	}
}

type bcdToStrTest struct {
	a string
	b []byte
}

var bcdTests = []bcdToStrTest{
	bcdToStrTest{"37", []byte{0x33, 0x37}},
} //37   02     12   04   25   7    00   00

func TestBCDToStr(t *testing.T) {
	for _, v := range bcdTests {
		ret := ByteToCtrlStr(v.b)
		if ret != v.a {
			t.Errorf("%v  want %v,but get %s", v.b, v.a, ret)
		}
	}

}

func TestGetCMTimeStrBySec(t *testing.T) {
	str := GetCMTimeStrBySec(time.Now().UnixNano() / 1e9)

	fmt.Println(str)

	str = GetCMTimeStrBySec(0)
	fmt.Println(str)
	fmt.Println(constdefine.INT32_MAX)
}

func TestIntToFloat(t *testing.T) {
	fmt.Println(TwoPrecision(402))

}

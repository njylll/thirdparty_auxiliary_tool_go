package xiamenprotocol

import (
	"fmt"
	"testing"
)

func TestCodec_UnMarshal(t *testing.T) {
	codec := &Codec{}
	req := &LoginReq{}
	req.DevAddr = "111111111111"
	req.GunNum = 2
	req.DevType = "99999"
	req.SoftVersion = "9999"
	req.VendorID = 2
	marshal, err := codec.Marshal(req)
	if err != nil {
		fmt.Println(err, marshal)
		t.Fail()
	}
	data := []byte{49, 49, 48, 49, 48, 49, 49, 49, 49, 48, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 0, 0, 0, 0, 0, 0, 69, 86, 68, 55, 53, 49, 95, 49, 50, 48, 75, 95, 84, 68, 0, 0, 0, 0, 0, 0, 0, 3, 86, 52, 46, 49, 48, 66, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 112, 23, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 0, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 0}
	reqAns := &LoginReq{}
	err = codec.UnMarshal(data, reqAns)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(reqAns.DevAddr)

}

func TestFrame_Marshal(t *testing.T) {
	codec := &Codec{}
	msgFrame := NewXiaMenProtocol()
	msgFrame.Cmd = 145
	msgFrame.Data = []byte{2, 20, 0, 0, 0, 0, 0, 0, 0}
	msgFrame.PacketID = 1
	msgFrame.FrameLength = uint16(len(msgFrame.Data) + 14 - 2)
	marshal, err := codec.Marshal(msgFrame)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(marshal)

}

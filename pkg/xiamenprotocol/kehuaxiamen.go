package xiamenprotocol

import (
	"bytes"
	"constdefine"
	"encoding/binary"
	"io"
	stob "struct2bytes"
	"thirdprotocol"
)

var (
	// ByteOrder 协议字节序
	ByteOrder = binary.LittleEndian
)

func init() {
	// 注册编解码器
	thirdprotocol.RegTransferProtocol(constdefine.XIAMEN, NewCodec)
}

// 厦门市政的帧定义
type XiaMenProtocol struct {
	Start1      byte   // 起始域
	Start2      byte   // 起始域
	FrameLength uint16 // 长度域 长度为2
	PacketID    byte   //计数器
	Source      uint32 //系统分配的每个桩的唯一 ID 号，0 表示后台系统。
	Dest        uint32
	Cmd         byte  //功能码，其中低 7 位表示功能码，最高位为 0 表示是请求命令，最高位为 1 表示是对该请求的应答。
	Data        Data  //功能码对应的数据域。长度为 Packet Length – 12B。
	CRC         int16 //前面数据的 CRC16 校验值。
}

func NewXiaMenProtocol() *XiaMenProtocol {
	x := &XiaMenProtocol{
		Start1: constdefine.XIAMENHEAD1,
		Start2: constdefine.XIAMENHEAD2,
	}
	return x
}

// Data 数据域类型
type Data []byte

func (d *Data) Dumps(buf *bytes.Buffer, ctx *stob.Context) error {
	_, err := buf.Write(*d)
	return err
}

func (d *Data) Loads(rdr *bytes.Reader, ctx *stob.Context) error {
	dLen := rdr.Len() - 2 // 剩余长度减去校验和的长度，就是数据域长度

	if len(*d) < dLen {
		if cap(*d) < dLen {
			*d = make(Data, dLen)
		} else {
			*d = (*d)[:dLen]
		}
	}
	n, _ := rdr.Read(*d)
	if n != dLen {
		return io.EOF
	}

	return nil
}

// 计算校验和
func (*XiaMenProtocol) AfterDumps(wire []byte) error {
	csBegin := len(wire) - 2
	csCalBegin := 0
	csCalEnd := len(wire) - 2
	binary.LittleEndian.PutUint16(wire[csBegin:csBegin+2], CRC16(wire[csCalBegin:csCalEnd]))
	return nil
}

//func ReadXiaMenProtocolFromBuff(data []byte) (*XiaMenProtocol, error) {
//
//	// 判断报文主体是否与PayloadLength相同
//	c := &XiaMenProtocol{
//		StartSymbol: data[0:2],
//		FrameLength: data[2:4],
//		PacketID:    data[4],
//		Source:      data[5:9],
//		Dest:        data[9:13],
//		Cmd:         data[13],
//	}
//	var dataLen int16
//	err := binary.Read(bytes.NewBuffer(c.FrameLength), ByteOrder, &dataLen)
//	if err != nil {
//		return nil, err
//	}
//	crcIndex := 13 + dataLen
//	c.Data = data[13:crcIndex]
//	c.CRC = data[crcIndex : crcIndex+2]
//	return c, nil
//}

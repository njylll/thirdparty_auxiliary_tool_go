/*
 * Created by 丁德鑫  on 2017/9/30 18:00:51
 * ==============================================
 * 包说明：通讯协议处理，主要处理封包和解包的过程
 * ==============================================
 * 南京德睿能源研究院版权所有
 * ==============================================
 */

package protoparse

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/constdefine"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/feslib"
	"github.com/njylll/thirdparty_auxiliary_tool_go/pkg/xiamenprotocol"
	"time"
)

var iv = "7645392627221362"
var key = "65ghf288sh28j234" //密文key

/**
*报文加密
*param encodeBytes 源数据
*param key	密钥
*return 加密后的报文
 */
func aesEncrypt(encodeBytes []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	encodeBytes = pKCS5Padding(encodeBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return crypted, nil
}
func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

// 根据报文头和长度寻找一帧，放到readerChannel中
func Unpack(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.Fixheaderlength { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.ConstHeader && buffer[i+1] == constdefine.Fixheaderlength { //0x63
			payloadLength := BytesToInt(buffer[i+constdefine.PayloadStartIndex : i+constdefine.PayloadStartIndex+constdefine.PayloadDataLength])
			// 可能为负值
			if payloadLength < 0 {
				feslib.Log(payloadLength)
				i = length
				break
			}
			if length < i+constdefine.Fixheaderlength+payloadLength { //如果长度不够
				break
			}
			data := buffer[i : i+constdefine.Fixheaderlength+payloadLength] //否则的话就取出这一帧
			*lastRecvTime = time.Now().Unix()                               //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {                                            //判断规约go程是否存在
				readerChannel <- data
				i += constdefine.Fixheaderlength + payloadLength - 1

				// 有可能越界
				if i < 0 {
					feslib.Log(i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据报文头和长度寻找一帧，放到readerChannel中
func UnpackChargeBow(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.ChargeBowFixheaderlength { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.ConstHeader && buffer[i+1] == constdefine.ChargeBowFixheaderlength { //0x63
			payloadLength := BytesToInt(buffer[i+constdefine.PayloadStartIndex : i+constdefine.PayloadStartIndex+constdefine.PayloadDataLength])
			// 可能为负值
			if payloadLength < 0 {
				feslib.Log(payloadLength)
				i = length
				break
			}
			if length < i+constdefine.ChargeBowFixheaderlength+payloadLength { //如果长度不够
				break
			}
			data := buffer[i : i+constdefine.ChargeBowFixheaderlength+payloadLength] //否则的话就取出这一帧
			*lastRecvTime = time.Now().Unix()                                        //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {                                                     //判断规约go程是否存在
				readerChannel <- data
				i += constdefine.ChargeBowFixheaderlength + payloadLength - 1

				// 有可能越界
				if i < 0 {
					feslib.Log(i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 发送集控器报文增加报文头     EncodingType: 消息体编码类型，Protobuf为0x01；JSON为0x02
func PacketToExtCtrl(msgType byte, ctrlAddrBytes []byte, message []byte, encodingType byte, usingaes int) []byte {
	head := [constdefine.Fixheaderlength]byte{0: constdefine.ConstHeader, 1: constdefine.Fixheaderlength, 2: constdefine.ProtocolType, 3: encodingType}

	var err error
	if constdefine.OK == usingaes {
		b64Byte := base64.StdEncoding.EncodeToString(message)
		message, err = aesEncrypt([]byte(b64Byte), []byte(key))
	}
	if err != nil {
		feslib.Log("AES encrypt error...")
	}

	payloadlen := feslib.IntToBytes(len(message))
	head[4] = payloadlen[0]
	head[5] = payloadlen[1]
	head[6] = payloadlen[2]
	head[7] = payloadlen[3]
	head[constdefine.MessageTypeIndex] = msgType
	head[9] = 0
	head[constdefine.ServerCodeIndex] = 1 //此处需要为1，因为message_copie中判断了是A模式的才会记录报文，否则可视化页面上看不到
	actualLen := len(ctrlAddrBytes)
	for i := 0; i < actualLen; i++ { //集控器地址
		head[constdefine.CtrlAddrIndex+i] = ctrlAddrBytes[i]
	}

	return append(head[:], message...)
}

// 发送集控器报文增加报文头
func PacketToCtrl(msgType byte, ctrlAddrBytes []byte, message []byte) []byte {
	head := [constdefine.Fixheaderlength]byte{0: constdefine.ConstHeader, 1: constdefine.Fixheaderlength, 2: constdefine.ProtocolType, 3: constdefine.EncodingType}

	payloadlen := feslib.IntToBytes(len(message))
	head[4] = payloadlen[0]
	head[5] = payloadlen[1]
	head[6] = payloadlen[2]
	head[7] = payloadlen[3]
	head[constdefine.MessageTypeIndex] = msgType
	head[9] = 0
	head[constdefine.ServerCodeIndex] = 0
	actualLen := len(ctrlAddrBytes)
	for i := 0; i < actualLen; i++ { //集控器地址
		head[constdefine.CtrlAddrIndex+i] = ctrlAddrBytes[i]
	}

	return append(head[:], message...)
}

// 发送充电弓报文增加报文头
func PacketToChargeBow(msgType byte, chargeBowAddrBytes []byte, message []byte) []byte {
	head := [constdefine.Fixheaderlength]byte{0: constdefine.ConstHeader, 1: constdefine.Fixheaderlength, 2: constdefine.ProtocolType, 3: constdefine.EncodingType}

	payloadlen := feslib.IntToBytes(len(message))
	head[4] = payloadlen[0]
	head[5] = payloadlen[1]
	head[6] = payloadlen[2]
	head[7] = payloadlen[3]
	head[constdefine.MessageTypeIndex] = msgType
	head[9] = 0
	head[constdefine.ServerCodeIndex] = 0
	actualLen := len(chargeBowAddrBytes)
	for i := 0; i < actualLen; i++ { //集控器地址
		head[constdefine.CtrlAddrIndex+i] = chargeBowAddrBytes[i]
	}

	return append(head[:], message...)
}

// 整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

// int64整形转换成8字节
func Int64ToBytes(n int64) []byte {
	x := int64(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()

}

// 根据蜗牛之家报文头和长度寻找一帧，放到readerChannel中
func UnpackSnail(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.SNAILHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.SNAILHEAD1 && buffer[i+1] == constdefine.SNAILHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.SNAILLENINDEX : i+constdefine.SNAILLENINDEX+constdefine.SNAILHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackSnail err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[constdefine.SNAILMSGINDEX : frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackSnail err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 积成电子解析
func UnPackIeslab(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		// 找到起始数据
		if buffer[i] == 0x7E {
			if endIndex := findIselabEnd(buffer, i); endIndex > 0 {
				data := buffer[i : endIndex+1]
				*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
				if *consumeIfExist {              //判断规约go程是否存在
					readerChannel <- data
					i = endIndex + 1
					// 有可能越界
					if i < 0 {
						feslib.Log("UnpackIeslab err i<0", i)
						i = length
						break
					}
				} else {
					feslib.Log("消费者退出了")
					break
				}
			} else {
				// 没有数据
				break
			}
		}
	}

	if i >= length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

func findIselabEnd(data []byte, startIndex int) int {
	startIndex++
	for ; startIndex < len(data); startIndex++ {
		if data[startIndex] == 0x7E {
			return startIndex
		}
	}

	return -1
}

// 绿能报文头和长度寻找一帧，放到readerChannel中
func UnpackLvNeng(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i++ {
		if length < i+constdefine.LVNENGHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.LVNENGHEAD1 && buffer[i+1] == constdefine.LVNENGHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.LVNENGENINDEX : i+constdefine.LVNENGENINDEX+constdefine.LVNENGHEADFRAMELEN]))
			if frameLen <= constdefine.LVNENGHEADLEN { // 判断帧长度是否过小
				continue
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[constdefine.LVNENGMSGINDEX : frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackLvNeng err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据盛弘报文头和长度寻找一帧，放到readerChannel中
func UnpackSinExcel(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.SINEXCELHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.SINEXCELHEAD1 && buffer[i+1] == constdefine.SINEXCELHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.SINEXCELLENINDEX : i+constdefine.SINEXCELLENINDEX+constdefine.SINEXCELHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackSinExcel err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[constdefine.SINEXCELMSGINDEX : frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackSinExcel err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据厦门报文头和长度寻找一帧，放到readerChannel中
func UnpackXiaMen(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.XIAMENHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.XIAMENHEAD1 && buffer[i+1] == constdefine.XIAMENHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.XIAMENLENINDEX : i+constdefine.XIAMENLENINDEX+constdefine.XIAMENHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackXiaMen err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen+4] //否则的话就取出这一帧

			calCheck := xiamenprotocol.CRC16(data[0 : frameLen+4-2])                // 计算的校验和
			getCheck := binary.LittleEndian.Uint16(data[frameLen+4-2 : frameLen+4]) // 报文中的校验和
			if calCheck != getCheck {
				feslib.Log("checksum XiaMen err cal", calCheck, "get", getCheck)
				continue
			}

			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen + 4 - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackXiaMen err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据橙电报文头和长度寻找一帧，放到readerChannel中
func UnpackChengDian(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.CHENGDIANHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.CHENGDIANHEAD1 && buffer[i+1] == constdefine.CHENGDIANHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.CHENGDIANLENINDEX : i+constdefine.CHENGDIANLENINDEX+constdefine.CHENGDIANHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackCHENGDIAN err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[constdefine.CHENGDIANMSGINDEX : frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackCHENGDIAN err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据新页报文头和长度寻找一帧，放到readerChannel中
func UnpackXinYe(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.XinYeHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.XinYeHEAD1 && buffer[i+1] == constdefine.XinYeHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.XinYeLENINDEX : i+constdefine.XinYeLENINDEX+constdefine.XinYeHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackSinExcel err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[constdefine.XinYeMSGINDEX : frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackSinExcel err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据永联报文头和长度寻找一帧，放到readerChannel中
func UnpackYongLian(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.YLHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.YLHEAD1 && buffer[i+1] == constdefine.YLHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.YLLENINDEX : i+constdefine.YLLENINDEX+constdefine.YLHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackYL err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[constdefine.YLMSGINDEX : frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackYL err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据盛弘报文头和长度寻找一帧，放到readerChannel中
func UnpackTeld(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.TELDHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.TELDHEAD1 && buffer[i+1] == constdefine.TELDHEAD2 {
			frameLen := int(feslib.ByteToUint(buffer[i+constdefine.TELDLENINDEX:i+constdefine.TELDLENINDEX+constdefine.TELDHEADFRAMELEN])) + constdefine.TELDHEADLEN + 1
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackTeld err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[:frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackTeld err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据星星报文头和长度寻找一帧，放到readerChannel中
func UnpackXingXing(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.XingXingHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.XingXingHEAD1 {
			frameLen := int(feslib.ByteToUint(buffer[i+constdefine.XingXingLENINDEX:i+constdefine.XingXingLENINDEX+constdefine.XingXingFRAMELEN])) + constdefine.XingXingHEADLEN + 1
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackXingXing err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[:frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackXingXing err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据万泊报文头和长度寻找一帧，放到readerChannel中
func UnpackEnPlus(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.ENPLUSHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.ENPLUSHEAD1 && buffer[i+1] == constdefine.ENPLUSHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.ENPLUSLENINDEX : i+constdefine.ENPLUSLENINDEX+constdefine.ENPLUSHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackEnPlus err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[constdefine.ENPLUSMSGINDEX : frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackEnPlus err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 累计和校验返回最后一个字节
func CheckSum(data []byte) byte {
	sum := 0
	i := 0
	for i = 0; i < len(data); i++ {
		sum += int(data[i])
	}
	return byte(sum & 0xff)
}

// 累计和校验返回最后一个字节
func GuoYaoCheckSum(data []byte) byte {
	sum := 0
	i := 0
	for i = 0; i < len(data); i++ {
		sum += int(data[i])
	}
	return byte(sum % 256)
}

// 蜗牛之家发送报文
func SnailPacketToCtrl(msgType int, seq byte, message []byte) []byte {
	frameLen := len(message) + constdefine.SNAILHEADLEN + 1
	frameLenBytes := feslib.IntTo2LillteBytes(frameLen)
	msgBytes := feslib.IntTo2LillteBytes(msgType)
	head := make([]byte, frameLen)
	head[0] = constdefine.SNAILHEAD1
	head[1] = constdefine.SNAILHEAD2
	head[2] = frameLenBytes[0]
	head[3] = frameLenBytes[1]
	head[4] = constdefine.SNAILVERSION
	head[5] = seq
	head[6] = msgBytes[0]
	head[7] = msgBytes[1]
	copy(head[constdefine.SNAILHEADLEN:], message) //payload
	head[frameLen-1] = CheckSum(head[constdefine.SNAILMSGINDEX : frameLen-1])
	return head[:]
}

// ///////科华桩协议////////////
func UnpackKehua(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.KEHUAHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.KEHUAHHEAD1 && buffer[i+1] == constdefine.KEHUAHHEAD2 {
			frameLen := int(feslib.ByteToUint16(buffer[i+constdefine.KEHUAHLENINDEX : i+constdefine.KEHUAHLENINDEX+constdefine.KEHUAHHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackKehua err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			checkBytes := feslib.IntTo2Bytes(int(CheckCRC16Sum(data[:frameLen-3])))
			if checkBytes[0] != data[frameLen-3] || checkBytes[1] != data[frameLen-2] {
				feslib.Log("UnpackKehua checksum err want:", data[frameLen-3:frameLen-1], "now", checkBytes)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackKehua  err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// ///////最新版本云快充协议////////////
func UnpackNewYunkc(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.NewYUNKCHEADLENGTH { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.NEWYUNKCHEAD {
			frameLen := int(buffer[i+constdefine.NEWYUNKCLENINDEX]) + 4

			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackNewYunkc err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//feslib.Log("解出一帧", data)
			checkBytes := feslib.IntTo2Bytes(UsMBCRC16(data[2:frameLen-2], frameLen-4))
			if checkBytes[0] != data[frameLen-2] || checkBytes[1] != data[frameLen-1] {
				feslib.Log("UnpackNewYUNKC checksum err want:", data[frameLen-2:frameLen], "now", checkBytes)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1
				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackNewYUNKC  err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}
	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// ///////云快充104协议////////////
func UnpackYunkc(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	frameLen := 0
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.YUNKCHEARTBEATLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.YUNKCHEAD && buffer[i+1] == constdefine.YUNKCPROTOTYPE && constdefine.YUNKCDEVTYPE == buffer[i+2] && length >= i+constdefine.YUNKCSTARTFIXEDHEADLEN { //协议启动帧
			//feslib.Log("协议启动帧")
			frameLen = constdefine.YUNKCSTARTFIXEDHEADLEN
		} else if buffer[i] == constdefine.YUNKCHEAD && buffer[i+1] == constdefine.YUNKCHEARTBEAT && (constdefine.YUNKCHEARTBEATThird == buffer[i+2] || 0x83 == buffer[i+2]) && constdefine.YUNKCHEARTBEATFour == buffer[i+3] && constdefine.YUNKCHEARTBEATFouth == buffer[i+4] && constdefine.YUNKCHEARTBEATSix == buffer[i+5] && length >= i+constdefine.YUNKCHEARTBEATLEN { //心跳报文需特殊处理
			frameLen = constdefine.YUNKCHEARTBEATLEN
		} else if buffer[i] == constdefine.YUNKCHEAD {
			frameLen = int(buffer[i+1]) + 2
		} else {
			continue
		}
		// 可能为负值
		if frameLen <= 0 {
			feslib.Log("UnpackYunkc err frameLen<0", frameLen)
			i = length
			break
		}
		if length < i+frameLen { //如果长度不够
			break
		}
		data := buffer[i : i+frameLen] //否则的话就取出这一帧
		//feslib.Log("解出一帧", data)

		*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
		if *consumeIfExist {              //判断规约go程是否存在
			readerChannel <- data
			i += frameLen - 1

			// 有可能越界
			if i < 0 {
				feslib.Log("UnpackYunkc  err i<0", i)
				i = length
				break
			}
		} else {
			feslib.Log("消费者退出了")
			break
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// ///////DongXu协议////////////
func UnpackDongXu(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.DONGXUHEADLENGTH { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.DONGXUHEAD {
			frameLen := int(buffer[i+constdefine.DONGXULENINDEX]) + 4

			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackDONGXU err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//feslib.Log("解出一帧", data)
			checkVal := UsMBCRC16(data[2:frameLen-2], frameLen-4) // 计算的校验和
			recvCsBytes := data[frameLen-2:]
			recvCsValBig := int(binary.BigEndian.Uint16(recvCsBytes))
			recvCsValLit := int(binary.LittleEndian.Uint16(recvCsBytes))
			if checkVal != recvCsValBig && checkVal != recvCsValLit { // * 适配人民电器，校验和正反序有一个能对上就行
				feslib.Log("UnpackDONGXU checksum err want:", recvCsBytes, "now", feslib.IntTo2Bytes(checkVal))
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1
				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackDONGXU  err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}
	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据国耀报文头和长度寻找一帧，放到readerChannel中
func UnpackGuoYao(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.GUOYAOHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}

		if buffer[i] == constdefine.GUOYAOHEAD1 {
			frameLen := constdefine.GUOYAOHEADLEN + int(feslib.ByteToUint16Little(buffer[i+constdefine.GUOYAOLENINDEX:i+constdefine.GUOYAOLENINDEX+constdefine.GUOYAOLHEADFRAMELEN])) + constdefine.GUOYAOENDLENGTH
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackGuoYao err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := GuoYaoCheckSum(data[:frameLen-2])
			if check != data[frameLen-2] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}

			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackGuoYao err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// ///////聚电104协议////////////
func UnpackCostar(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	frameLen := 0
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.COSTARHEARTBEATLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.COSTARHEAD && buffer[i+1] == constdefine.COSTARPROTOTYPE && constdefine.COSTARDEVTYPE == buffer[i+2] && length >= i+constdefine.COSTARSTARTFIXEDHEADLEN { //协议启动帧
			//feslib.Log("协议启动帧")
			frameLen = constdefine.COSTARSTARTFIXEDHEADLEN
		} else if buffer[i] == constdefine.COSTARHEAD && buffer[i+1] == constdefine.COSTARHEARTBEAT && (constdefine.COSTARHEARTBEATThird == buffer[i+2] || 0x83 == buffer[i+2]) && constdefine.COSTARHEARTBEATFour == buffer[i+3] && constdefine.COSTARHEARTBEATFouth == buffer[i+4] && constdefine.COSTARHEARTBEATSix == buffer[i+5] && length >= i+constdefine.COSTARHEARTBEATLEN { //心跳报文需特殊处理
			frameLen = constdefine.COSTARHEARTBEATLEN
		} else if buffer[i] == constdefine.COSTARHEAD && buffer[i+1] == constdefine.COSTARHEARTBEAT && (0x07 == buffer[i+2] || 0x0b == buffer[i+2]) && constdefine.COSTARHEARTBEATFour == buffer[i+3] && constdefine.COSTARHEARTBEATFouth == buffer[i+4] && constdefine.COSTARHEARTBEATSix == buffer[i+5] && length >= i+constdefine.COSTARHEARTBEATLEN { //心跳报文需特殊处理{
			frameLen = constdefine.COSTARHEARTBEATLEN
		} else if buffer[i] == constdefine.COSTARHEAD {
			frameLen = int(buffer[i+1]) + 2
		} else {
			continue
		}
		// 可能为负值
		if frameLen <= 0 {
			feslib.Log("UnpackYunkc err frameLen<0", frameLen)
			i = length
			break
		}
		if length < i+frameLen { //如果长度不够
			break
		}
		data := buffer[i : i+frameLen] //否则的话就取出这一帧
		//feslib.Log("解出一帧", data)

		*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
		if *consumeIfExist {              //判断规约go程是否存在
			readerChannel <- data
			i += frameLen - 1

			// 有可能越界
			if i < 0 {
				feslib.Log("UnpackYunkc  err i<0", i)
				i = length
				break
			}
		} else {
			feslib.Log("消费者退出了")
			break
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 根据蜗牛之家报文头和长度寻找一帧，放到readerChannel中
func UnpackSse(buffer []byte, readerChannel chan []byte, lastRecvTime *int64, consumeIfExist *bool) []byte {
	length := len(buffer)
	tmpBuffer := make([]byte, 0)
	var i int
	for i = 0; i < length; i = i + 1 {
		if length < i+constdefine.SSEHEADLEN { //如果这一包数据长度小于报文头+报文长度
			break
		}
		if buffer[i] == constdefine.SSEHEAD1 && buffer[i+1] == constdefine.SSEHEAD2 {
			frameLen := int(feslib.ByteToUint16Little(buffer[i+constdefine.SSELENINDEX : i+constdefine.SSELENINDEX+constdefine.SSEHEADFRAMELEN]))
			// 可能为负值
			if frameLen <= 0 {
				feslib.Log("UnpackSSE err frameLen<0", frameLen)
				i = length
				break
			}
			if length < i+frameLen { //如果长度不够
				break
			}
			data := buffer[i : i+frameLen] //否则的话就取出这一帧
			//debug feslib.Log("解出一帧", data)
			check := CheckSum(data[0 : frameLen-1])
			if check != data[frameLen-1] {
				feslib.Log("checksum err want:", data[frameLen-1], "now", check)
			}
			*lastRecvTime = time.Now().Unix() //取出一帧就更新心跳超时的判断时间
			if *consumeIfExist {              //判断规约go程是否存在
				readerChannel <- data
				i += frameLen - 1

				// 有可能越界
				if i < 0 {
					feslib.Log("UnpackSSE err i<0", i)
					i = length
					break
				}
			} else {
				feslib.Log("消费者退出了")
				break
			}
		}
	}

	if i == length { //如果都遍历结束了
		return tmpBuffer //make([]byte, 0)
	}
	return buffer[i:] //否则不够的话，缓存下这些数据，等着收到后面的一起再处理
}

// 云快充CRC 计算函数
var (
	aucCRCHi = []byte{
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
		0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
		0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
		0x00, 0xC1, 0x81, 0x40,
	}

	aucCRCLo = []byte{
		0x00, 0xC0, 0xC1, 0x01, 0xC3, 0x03, 0x02, 0xC2, 0xC6, 0x06, 0x07, 0xC7,
		0x05, 0xC5, 0xC4, 0x04, 0xCC, 0x0C, 0x0D, 0xCD, 0x0F, 0xCF, 0xCE, 0x0E,
		0x0A, 0xCA, 0xCB, 0x0B, 0xC9, 0x09, 0x08, 0xC8, 0xD8, 0x18, 0x19, 0xD9,
		0x1B, 0xDB, 0xDA, 0x1A, 0x1E, 0xDE, 0xDF, 0x1F, 0xDD, 0x1D, 0x1C, 0xDC,
		0x14, 0xD4, 0xD5, 0x15, 0xD7, 0x17, 0x16, 0xD6, 0xD2, 0x12, 0x13, 0xD3,
		0x11, 0xD1, 0xD0, 0x10, 0xF0, 0x30, 0x31, 0xF1, 0x33, 0xF3, 0xF2, 0x32,
		0x36, 0xF6, 0xF7, 0x37, 0xF5, 0x35, 0x34, 0xF4, 0x3C, 0xFC, 0xFD, 0x3D,
		0xFF, 0x3F, 0x3E, 0xFE, 0xFA, 0x3A, 0x3B, 0xFB, 0x39, 0xF9, 0xF8, 0x38,
		0x28, 0xE8, 0xE9, 0x29, 0xEB, 0x2B, 0x2A, 0xEA, 0xEE, 0x2E, 0x2F, 0xEF,
		0x2D, 0xED, 0xEC, 0x2C, 0xE4, 0x24, 0x25, 0xE5, 0x27, 0xE7, 0xE6, 0x26,
		0x22, 0xE2, 0xE3, 0x23, 0xE1, 0x21, 0x20, 0xE0, 0xA0, 0x60, 0x61, 0xA1,
		0x63, 0xA3, 0xA2, 0x62, 0x66, 0xA6, 0xA7, 0x67, 0xA5, 0x65, 0x64, 0xA4,
		0x6C, 0xAC, 0xAD, 0x6D, 0xAF, 0x6F, 0x6E, 0xAE, 0xAA, 0x6A, 0x6B, 0xAB,
		0x69, 0xA9, 0xA8, 0x68, 0x78, 0xB8, 0xB9, 0x79, 0xBB, 0x7B, 0x7A, 0xBA,
		0xBE, 0x7E, 0x7F, 0xBF, 0x7D, 0xBD, 0xBC, 0x7C, 0xB4, 0x74, 0x75, 0xB5,
		0x77, 0xB7, 0xB6, 0x76, 0x72, 0xB2, 0xB3, 0x73, 0xB1, 0x71, 0x70, 0xB0,
		0x50, 0x90, 0x91, 0x51, 0x93, 0x53, 0x52, 0x92, 0x96, 0x56, 0x57, 0x97,
		0x55, 0x95, 0x94, 0x54, 0x9C, 0x5C, 0x5D, 0x9D, 0x5F, 0x9F, 0x9E, 0x5E,
		0x5A, 0x9A, 0x9B, 0x5B, 0x99, 0x59, 0x58, 0x98, 0x88, 0x48, 0x49, 0x89,
		0x4B, 0x8B, 0x8A, 0x4A, 0x4E, 0x8E, 0x8F, 0x4F, 0x8D, 0x4D, 0x4C, 0x8C,
		0x44, 0x84, 0x85, 0x45, 0x87, 0x47, 0x46, 0x86, 0x82, 0x42, 0x43, 0x83,
		0x41, 0x81, 0x80, 0x40,
	}
)

func UsMBCRC16(pucFrame []byte, usLen int) int {

	ucCRCHi := 0xFF

	ucCRCLo := 0xFF

	iIndex := 0

	for i := 0; i < usLen; i++ {

		iIndex = ucCRCLo ^ int(pucFrame[i])

		ucCRCLo = ucCRCHi ^ int(aucCRCHi[iIndex])

		ucCRCHi = int(aucCRCLo[iIndex])

	}
	return ucCRCHi<<8 | ucCRCLo
}

// CRC16
var MbTable = []uint16{
	0x0000, 0xC0C1, 0xC181, 0x0140, 0xC301, 0x03C0, 0x0280, 0xC241,
	0xC601, 0x06C0, 0x0780, 0xC741, 0x0500, 0xC5C1, 0xC481, 0x0440,
	0xCC01, 0x0CC0, 0x0D80, 0xCD41, 0x0F00, 0xCFC1, 0xCE81, 0x0E40,
	0x0A00, 0xCAC1, 0xCB81, 0x0B40, 0xC901, 0x09C0, 0x0880, 0xC841,
	0xD801, 0x18C0, 0x1980, 0xD941, 0x1B00, 0xDBC1, 0xDA81, 0x1A40,
	0x1E00, 0xDEC1, 0xDF81, 0x1F40, 0xDD01, 0x1DC0, 0x1C80, 0xDC41,
	0x1400, 0xD4C1, 0xD581, 0x1540, 0xD701, 0x17C0, 0x1680, 0xD641,
	0xD201, 0x12C0, 0x1380, 0xD341, 0x1100, 0xD1C1, 0xD081, 0x1040,
	0xF001, 0x30C0, 0x3180, 0xF141, 0x3300, 0xF3C1, 0xF281, 0x3240,
	0x3600, 0xF6C1, 0xF781, 0x3740, 0xF501, 0x35C0, 0x3480, 0xF441,
	0x3C00, 0xFCC1, 0xFD81, 0x3D40, 0xFF01, 0x3FC0, 0x3E80, 0xFE41,
	0xFA01, 0x3AC0, 0x3B80, 0xFB41, 0x3900, 0xF9C1, 0xF881, 0x3840,
	0x2800, 0xE8C1, 0xE981, 0x2940, 0xEB01, 0x2BC0, 0x2A80, 0xEA41,
	0xEE01, 0x2EC0, 0x2F80, 0xEF41, 0x2D00, 0xEDC1, 0xEC81, 0x2C40,
	0xE401, 0x24C0, 0x2580, 0xE541, 0x2700, 0xE7C1, 0xE681, 0x2640,
	0x2200, 0xE2C1, 0xE381, 0x2340, 0xE101, 0x21C0, 0x2080, 0xE041,
	0xA001, 0x60C0, 0x6180, 0xA141, 0x6300, 0xA3C1, 0xA281, 0x6240,
	0x6600, 0xA6C1, 0xA781, 0x6740, 0xA501, 0x65C0, 0x6480, 0xA441,
	0x6C00, 0xACC1, 0xAD81, 0x6D40, 0xAF01, 0x6FC0, 0x6E80, 0xAE41,
	0xAA01, 0x6AC0, 0x6B80, 0xAB41, 0x6900, 0xA9C1, 0xA881, 0x6840,
	0x7800, 0xB8C1, 0xB981, 0x7940, 0xBB01, 0x7BC0, 0x7A80, 0xBA41,
	0xBE01, 0x7EC0, 0x7F80, 0xBF41, 0x7D00, 0xBDC1, 0xBC81, 0x7C40,
	0xB401, 0x74C0, 0x7580, 0xB541, 0x7700, 0xB7C1, 0xB681, 0x7640,
	0x7200, 0xB2C1, 0xB381, 0x7340, 0xB101, 0x71C0, 0x7080, 0xB041,
	0x5000, 0x90C1, 0x9181, 0x5140, 0x9301, 0x53C0, 0x5280, 0x9241,
	0x9601, 0x56C0, 0x5780, 0x9741, 0x5500, 0x95C1, 0x9481, 0x5440,
	0x9C01, 0x5CC0, 0x5D80, 0x9D41, 0x5F00, 0x9FC1, 0x9E81, 0x5E40,
	0x5A00, 0x9AC1, 0x9B81, 0x5B40, 0x9901, 0x59C0, 0x5880, 0x9841,
	0x8801, 0x48C0, 0x4980, 0x8941, 0x4B00, 0x8BC1, 0x8A81, 0x4A40,
	0x4E00, 0x8EC1, 0x8F81, 0x4F40, 0x8D01, 0x4DC0, 0x4C80, 0x8C41,
	0x4400, 0x84C1, 0x8581, 0x4540, 0x8701, 0x47C0, 0x4680, 0x8641,
	0x8201, 0x42C0, 0x4380, 0x8341, 0x4100, 0x81C1, 0x8081, 0x4040}

// func CheckCRC16Sum(data []byte) int16 {
func CheckCRC16Sum(data []byte) uint16 {
	var crc16 uint16
	crc16 = 0xffff
	for _, v := range data {
		n := uint8(uint16(v) ^ crc16)
		crc16 >>= 8
		crc16 ^= MbTable[n]
	}
	return crc16
	/*sum := 0
	i := 0
	for i = 0; i < len(data); i++ {
		sum += int(data[i])
	}
	return int16(sum & 0xFFFF)*/
}

// 送报文
func KehuaPacketToCtrl(devType byte, msgType int, seq int, SN []byte, gunNum byte, message []byte) []byte {
	frameLen := len(message) + constdefine.KEHUAHEADLEN + 3
	frameLenBytes := feslib.IntTo2Bytes(frameLen)
	seqBytes := feslib.IntTo2Bytes(seq)

	head := make([]byte, frameLen)
	head[0] = constdefine.KEHUAHHEAD1
	head[1] = constdefine.KEHUAHHEAD2
	head[2] = frameLenBytes[0]
	head[3] = frameLenBytes[1]
	head[4] = 0 //TODO 协议版本
	head[5] = 0

	head[6] = seqBytes[0]
	head[7] = seqBytes[1]
	head[8] = devType

	copy(head[9:], SN) //设备串号
	head[30] = byte(msgType)

	head[31] = 0
	head[32] = 1

	head[33] = gunNum
	copy(head[constdefine.KEHUAHEADLEN:], message) //payload

	checkBytes := feslib.IntTo2Bytes(int(CheckCRC16Sum(head[:frameLen-3])))
	head[frameLen-3] = checkBytes[0]
	head[frameLen-2] = checkBytes[1]
	head[frameLen-1] = 0x68

	return head[:]
}

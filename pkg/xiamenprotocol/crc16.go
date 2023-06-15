package xiamenprotocol

const (
	_TABLE_LEN = 256    // crc16查找表长度
	_BYTE_BITS = 8      // 一个字节的位数
	_CRC16_REV = 0xA001 // crc16多项式倒序
	_BIT_0     = 0x0001 // 用于获取最低位（LSB）
)

type crc16Table = [_TABLE_LEN]uint16

var table crc16Table

// 在初始化函数中初始化crc16查找表
func init() {
	initCRCTable(&table)
}

// 计算crc16查找表
func initCRCTable(t *crc16Table) {
	var lsb, crc uint16

	for i := uint16(0); i < _TABLE_LEN; i++ {
		crc = i
		for j := 0; j < _BYTE_BITS; j++ {
			lsb = crc & _BIT_0
			crc >>= 1
			if lsb > 0 {
				crc ^= _CRC16_REV
			}
		}
		t[i] = crc
	}
}

// CRC16 计算数据的crc16校验码
func CRC16(data []byte) uint16 {
	var crc16 uint16 = 0xffff
	for _, d := range data {
		idx := byte(crc16) ^ d
		crc16 = (crc16 >> _BYTE_BITS) ^ table[idx]
	}
	return crc16
}

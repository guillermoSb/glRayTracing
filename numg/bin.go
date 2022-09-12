package numg

import "encoding/binary"

func Word(value uint16) []byte {
	bs := make([]byte,2)
	binary.LittleEndian.PutUint16(bs,value)
	return bs
}

func Dword(value uint32) []byte {
	bs := make([]byte,4)
	binary.LittleEndian.PutUint32(bs,value)
	return bs
}
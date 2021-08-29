package util

import (
	"encoding/binary"
)

func ByteToInt(b byte) int {
	return int(b)
}

func IntToByte(n int) byte {
	return byte(n)
}

func EncodeInt(n int) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(n))
	return b
}

func EncodeInt64(n int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(n))
	return b
}

func EncodeArrayInt(array []int) [][]byte {
	ret := [][]byte{}
	for _, e := range array {
		ret = append(ret, EncodeInt(e))
	}
	return ret
}

func ArrayInt64ToInt64(array []int64) int64 {
	var total int64 = 0
	for _, v := range array {
		total += v
	}
	return total
}

func DecodeInt(d []byte) int {
	return int(int32(binary.LittleEndian.Uint32(d)))
}

func DecodeInt64(d []byte) int64 {
	return int64(binary.LittleEndian.Uint64(d))
}

func DecodeArrayInt(array [][]byte) []int {
	ret := []int{}
	for _, e := range array {
		ret = append(ret, DecodeInt(e))
	}
	return ret
}

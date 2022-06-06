package utils

import "strconv"

func Byte32ToByteSlice(b32 [32]byte) []byte {
	var bytes []byte
	for i := 0; i < 32; i++ {
		bytes = append(bytes, b32[i])
	}
	return bytes
}

func ByteSliceToByte32(byteSlice []byte) [32]byte {
	var bytes32 [32]byte
	for i := 0; i < len(byteSlice); i++ {
		bytes32[i] = byteSlice[i]
	}
	return bytes32
}

func StringToInt64(s string) int64 {
	int64Num, _ := strconv.ParseInt(s, 10, 64)
	return int64Num
}

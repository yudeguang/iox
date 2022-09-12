package iox

import "encoding/binary"

//int8ToBytes
func int8ToBytes(i int8) []byte {
	b := make([]byte, 1)
	b[0] = uint8(i)
	return b
}

//uint8ToBytes
func uint8ToBytes(i uint8) []byte {
	b := make([]byte, 1)
	b[0] = i
	return b
}

//int16ToBytes
func int16ToBytes(i int16) []byte {
	var buf = make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(i))
	return buf
}

//int16ToBytesBigEndian
func int16ToBytesBigEndian(i int16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i))
	return buf
}

//uint16ToBytes
func uint16ToBytes(i uint16) []byte {
	var buf = make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, i)
	return buf
}

//uint16ToBytesBigEndian
func uint16ToBytesBigEndian(i uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, i)
	return buf
}

//int32ToBytes
func int32ToBytes(i int32) []byte {
	var buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}

//int32ToBytesBigEndian
func int32ToBytesBigEndian(i int32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

//uint32ToBytes
func uint32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	return buf
}

//uint32ToBytesBigEndian
func uint32ToBytesBigEndian(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, i)
	return buf
}

//int64ToBytes
func int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

//int64ToBytesBigEndian
func int64ToBytesBigEndian(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

//uint64ToBytes
func uint64ToBytes(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, i)
	return buf
}

//uint64ToBytesBigEndian
func uint64ToBytesBigEndian(i uint64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

//bytesToInt8
func bytesToInt8(b []byte) int8 {
	return int8(b[0])
}

//bytesToUint8
func bytesToUint8(b []byte) uint8 {
	return b[0]
}

//bytesToInt16
func bytesToInt16(b []byte) int16 {
	return int16(binary.LittleEndian.Uint16(b))
}

//bytesToInt16BigEndian
func bytesToInt16BigEndian(b []byte) int16 {
	return int16(binary.BigEndian.Uint16(b))
}

//bytesToUint16
func bytesToUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

//bytesToUint16BigEndian
func bytesToUint16BigEndian(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

//bytesToInt32
func bytesToInt32(b []byte) int32 {
	return int32(binary.LittleEndian.Uint32(b))
}

//bytesToInt32BigEndian
func bytesToInt32BigEndian(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(b))
}

//bytesToUint32
func bytesToUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

//bytesToUint32BigEndian
func bytesToUint32BigEndian(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

//bytesToInt64
func bytesToInt64(b []byte) int64 {
	return int64(binary.LittleEndian.Uint64(b))
}

//bytesToInt64BigEndian
func bytesToInt64BigEndian(b []byte) int64 {
	return int64(binary.BigEndian.Uint64(b))
}

//bytesToUint64
func bytesToUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}

//bytesToUint64BigEndian
func bytesToUint64BigEndian(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
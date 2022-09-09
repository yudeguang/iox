package iox

import (
	"bytes"
	"strconv"
)

type Writer struct {
	writer bytes.Buffer
}

//初始化
func NewBytesBuffer(buf bytes.Buffer) *Writer {
	r := new(Writer)
	r.writer = buf
	return r
}

//输出
func (w *Writer) Bytes() []byte {
	return w.writer.Bytes()
}

//重置
func (w *Writer) Reset() {
	w.writer.Reset()
}

//写入字节
func (w *Writer) WriteByte(p []byte) {
	w.writer.Write(p)
}

//写入文本
func (w *Writer) WriteString(s string) {
	w.writer.Write([]byte(s))
}

//先写入字节数组的长度，然后再写入该字节数组，该字节数组长度必须是在unit8范围，否则抛出异常
func (w *Writer) WriteByteUint8(p []byte) {
	if int(uint8(len(p))) != len(p) {
		panic("the data length:" + strconv.Itoa(len(p)) + " is too big for Uint8")
	}
	w.WriteUint8(uint8(len(p)))
	w.writer.Write(p)
}

//先写入文本的长度，然后再写入该文本，该文本长度必须是在unit8范围，否则抛出异常
func (w *Writer) WriteStringUint8(s string) {
	w.WriteByteUint8([]byte(s))
}

//先写入字节数组的长度，然后再写入该字节数组，该字节数组长度必须是在unit16范围，否则抛出异常
func (w *Writer) WriteByteUint16(p []byte) {
	if int(uint16(len(p))) != len(p) {
		panic("the data length:" + strconv.Itoa(len(p)) + " is too big for Uint16")
	}
	w.WriteUint16(uint16(len(p)))
	w.writer.Write(p)
}

//先写入文本的长度，然后再写入该文本，该文本长度必须是在unit16范围，否则抛出异常
func (w *Writer) WriteStringUint16(s string) {
	w.WriteByteUint16([]byte(s))
}

//先写入字节数组的长度，然后再写入该字节数组，该字节数组长度必须是在unit32范围，否则抛出异常
func (w *Writer) WriteByteUint32(p []byte) {
	if int(uint32(len(p))) != len(p) {
		panic("the data length:" + strconv.Itoa(len(p)) + " is too big for Uint32")
	}
	w.WriteUint32(uint32(len(p)))
	w.writer.Write(p)
}

//先写入文本的长度，然后再写入该文本，该文本长度必须是在unit32范围，否则抛出异常
func (w *Writer) WriteStringUint32(s string) {
	w.WriteByteUint32([]byte(s))
}

//先写入文本的长度，然后再写入该文本，该文本长度必须是在Uint648范围，否则抛出异常
func (w *Writer) WriteStringUint64(s string) {
	w.WriteByteUint64([]byte(s))
}

//先写入字节数组的长度，然后再写入该字节数组，该字节数组长度必须是在unit64范围，否则抛出异常
func (w *Writer) WriteByteUint64(p []byte) {
	w.WriteUint64(uint64(len(p)))
	w.writer.Write(p)
}

//写入int18 无大端小端区别
func (w *Writer) WriteInt8(i int8) {
	w.writer.Write([]byte{uint8(i)})
}

//写入Uint8 无大端小端区别
func (w *Writer) WriteUint8(i uint8) {
	w.writer.Write([]byte{i})
}

//写入int16
func (w *Writer) WriteInt16(i int16) {
	w.writer.Write(int16ToBytes(i))
}

//写入int16 大端模式
func (w *Writer) WriteInt16BigEndian(i int16) {
	w.writer.Write(int16ToBytesBigEndian(i))
}

//写入uint16
func (w *Writer) WriteUint16(i uint16) {
	w.writer.Write(uint16ToBytes(i))
}

//写入uint16 大端模式
func (w *Writer) WriteUint16BigEndian(i uint16) {
	w.writer.Write(uint16ToBytesBigEndian(i))
}

//写入int32
func (w *Writer) WriteInt32(i int32) {
	w.writer.Write(int32ToBytes(i))
}

//写入int32 大端模式
func (w *Writer) WriteInt32BigEndian(i int32) {
	w.writer.Write(int32ToBytesBigEndian(i))
}

//写入uint32
func (w *Writer) WriteUint32(i uint32) {
	w.writer.Write(uint32ToBytes(i))
}

//写入uint32 大端模式
func (w *Writer) WriteUint32BigEndian(i uint32) {
	w.writer.Write(uint32ToBytesBigEndian(i))
}

//写入int64
func (w *Writer) WriteInt64(i int64) {
	w.writer.Write(int64ToBytes(i))
}

//写入 大端模式
func (w *Writer) WriteInt64BigEndian(i int64) {
	w.writer.Write(int64ToBytesBigEndian(i))
}

//写入uint64
func (w *Writer) WriteUint64(i uint64) {
	w.writer.Write(uint64ToBytes(i))
}

//写入uint64 大端模式
func (w *Writer) WriteUint64BigEndian(i uint64) {
	w.writer.Write(uint64ToBytesBigEndian(i))
}

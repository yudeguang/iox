package iox

import (
	"bytes"
	"math"
	"strconv"
)

// Writer helps you write data into an bytes.Buffer.
// the default ByteOrder is LittleEndian.
type Writer struct {
	writer bytes.Buffer
}

//NewBytesBuffer returns a *Writer.
func NewBytesBuffer(buf ...bytes.Buffer) *Writer {
	r := new(Writer)
	if len(buf) == 0 {
		var b bytes.Buffer
		r.writer = b
	} else {
		r.writer = buf[0]
	}
	return r
}

//when write finished get the data from the bytes.Buffer.
func (w *Writer) Bytes() []byte {
	return w.writer.Bytes()
}

//resets the buffer to be empty.
func (w *Writer) Reset() {
	w.writer.Reset()
}

//Write Byte into writer.
func (w *Writer) WriteBytes(p []byte) {
	w.writer.Write(p)
}

//Write String into writer
func (w *Writer) WriteString(s string) {
	w.writer.Write([]byte(s))
}

//Write the length(Uint8) of the byte first, then write the byte.
func (w *Writer) WriteBytesUint8(p []byte) {
	if int(uint8(len(p))) != len(p) {
		panic("the data length:" + strconv.Itoa(len(p)) + " is too big for Uint8")
	}
	w.WriteUint8(uint8(len(p)))
	w.writer.Write(p)
}

//Write the length(Uint8) of the string first, then write the string.
func (w *Writer) WriteStringUint8(s string) {
	w.WriteBytesUint8([]byte(s))
}

//Write the length(Uint16) of the byte first, then write the byte.
func (w *Writer) WriteBytesUint16(p []byte) {
	if int(uint16(len(p))) != len(p) {
		panic("the data length:" + strconv.Itoa(len(p)) + " is too big for Uint16")
	}
	w.WriteUint16(uint16(len(p)))
	w.writer.Write(p)
}

//Write the length(Uint16 BigEndian) of the byte first, then write the byte.
func (w *Writer) WriteBytesUint16BigEndian(p []byte) {
	if int(uint16(len(p))) != len(p) {
		panic("the data length:" + strconv.Itoa(len(p)) + " is too big for Uint16")
	}
	w.WriteUint16BigEndian(uint16(len(p)))
	w.writer.Write(p)
}

//Write the length(Uint16) of the string first, then write the string.
func (w *Writer) WriteStringUint16(s string) {
	w.WriteBytesUint16([]byte(s))
}

//Write the length(Uint16 BigEndian) of the string first, then write the string.
func (w *Writer) WriteStringUint16BigEndian(s string) {
	w.WriteBytesUint16BigEndian([]byte(s))
}

//Write the length(Uint32) of the byte first, then write the byte.
func (w *Writer) WriteBytesUint32(p []byte) {
	if int(uint32(len(p))) != len(p) {
		panic("the data length:" + strconv.Itoa(len(p)) + " is too big for Uint32")
	}
	w.WriteUint32(uint32(len(p)))
	w.writer.Write(p)
}

//Write the length(Uint32 BigEndian) of the byte first, then write the byte.
func (w *Writer) WriteBytesUint32BigEndian(p []byte) {
	if int(uint32(len(p))) != len(p) {
		panic("the data length:" + strconv.Itoa(len(p)) + " is too big for Uint32")
	}
	w.WriteUint32BigEndian(uint32(len(p)))
	w.writer.Write(p)
}

//Write the length(Uint32) of the string first, then write the string.
func (w *Writer) WriteStringUint32(s string) {
	w.WriteBytesUint32([]byte(s))
}

//Write the length(Uint32 BigEndian) of the string first, then write the string.
func (w *Writer) WriteStringUint32BigEndian(s string) {
	w.WriteBytesUint32BigEndian([]byte(s))
}

//Write the length(Uint16) of the byte first, then write the byte.
func (w *Writer) WriteBytesUint64(p []byte) {
	w.WriteUint64(uint64(len(p)))
	w.writer.Write(p)
}

//Write the length(Uint16) of the byte first, then write the byte.
func (w *Writer) WriteBytesUint64BigEndian(p []byte) {
	w.WriteUint64BigEndian(uint64(len(p)))
	w.writer.Write(p)
}

//Write the length(Uint16) of the string first, then write the string.
func (w *Writer) WriteStringUint64(s string) {
	w.WriteBytesUint64([]byte(s))
}

//Write the length(Uint16) of the string first, then write the string.
func (w *Writer) WriteStringUint64BigEndian(s string) {
	w.WriteBytesUint64BigEndian([]byte(s))
}

//Write int8 into Writer.
func (w *Writer) WriteInt8(i int8) {
	w.writer.Write([]byte{uint8(i)})
}

//Write uint8 into Writer.
func (w *Writer) WriteUint8(i uint8) {
	w.writer.Write([]byte{i})
}

//Write int16 with LittleEndian into Writer.
func (w *Writer) WriteInt16(i int16) {
	w.writer.Write(int16ToBytes(i))
}

//Write int16 with BigEndian into Writer.
func (w *Writer) WriteInt16BigEndian(i int16) {
	w.writer.Write(int16ToBytesBigEndian(i))
}

//Write uint16 with LittleEndian into Writer.
func (w *Writer) WriteUint16(i uint16) {
	w.writer.Write(uint16ToBytes(i))
}

//Write uint16 with BigEndian into Writer.
func (w *Writer) WriteUint16BigEndian(i uint16) {
	w.writer.Write(uint16ToBytesBigEndian(i))
}

//Write int32 with LittleEndian into Writer.
func (w *Writer) WriteInt32(i int32) {
	w.writer.Write(int32ToBytes(i))
}

//Write int32 with BigEndian into Writer.
func (w *Writer) WriteInt32BigEndian(i int32) {
	w.writer.Write(int32ToBytesBigEndian(i))
}

//Write uint32 with LittleEndian into Writer.
func (w *Writer) WriteUint32(i uint32) {
	w.writer.Write(uint32ToBytes(i))
}

//Write uint32 with BigEndian into Writer.
func (w *Writer) WriteUint32BigEndian(i uint32) {
	w.writer.Write(uint32ToBytesBigEndian(i))
}

//Write int64 with LittleEndian into Writer.
func (w *Writer) WriteInt64(i int64) {
	w.writer.Write(int64ToBytes(i))
}

//Write int64 with BigEndian into Writer.
func (w *Writer) WriteInt64BigEndian(i int64) {
	w.writer.Write(int64ToBytesBigEndian(i))
}

//Write uint64 with LittleEndian into Writer.
func (w *Writer) WriteUint64(i uint64) {
	w.writer.Write(uint64ToBytes(i))
}

//Write uint64 with BigEndian into Writer.
func (w *Writer) WriteUint64BigEndian(i uint64) {
	w.writer.Write(uint64ToBytesBigEndian(i))
}

//Write float32 with LittleEndian into Writer.
func (w *Writer) WriteFloat32(i float32) {
	w.WriteUint32(math.Float32bits(i))
}

//Write float32 with BigEndian into Writer.
func (w *Writer) WriteFloat32BigEndian(i float32) {
	w.WriteUint32BigEndian(math.Float32bits(i))
}

//Write float64 with LittleEndian into Writer.
func (w *Writer) WriteFloat64(i float64) {
	w.WriteUint64(math.Float64bits(i))
}

//Write float64 with BigEndian into Writer.
func (w *Writer) WriteFloat64BigEndian(i float64) {
	w.WriteUint64BigEndian(math.Float64bits(i))
}

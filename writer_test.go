package iox

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	var tint8 int8 = -100
	//int16ToBytes
	b := int8ToBytes(tint8)
	tint82 := bytesToInt8(b)
	if tint8 != tint82 {
		t.Fatalf("unexpected value obtained; got %b want %b", tint8, tint82)
	}
	var tuint8 uint8 = 100
	//int16ToBytes
	b = uint8ToBytes(tuint8)
	tuint82 := bytesToUint8(b)
	if tuint82 != tuint82 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint8, tuint82)
	}

	var tint16 int16 = -1000
	//int16ToBytes
	b = int16ToBytes(tint16)
	tInt162 := bytesToInt16(b)
	if tint16 != tInt162 {
		t.Fatalf("unexpected value obtained; got %b want %b", tint16, tInt162)
	}
	//int16ToBytesBigEndian
	b = int16ToBytesBigEndian(tint16)
	tInt162 = bytesToInt16BigEndian(b)
	if tint16 != tInt162 {
		t.Fatalf("unexpected value obtained; got %b want %b", tint16, tInt162)
	}

	var tuint16 uint16 = 1000
	//uint16ToBytes
	b = uint16ToBytes(tuint16)
	tuint162 := bytesToUint16(b)
	if tuint16 != tuint162 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint16, tuint162)
	}
	b = uint16ToBytesBigEndian(tuint16)
	tuint162 = bytesToUint16BigEndian(b)
	if tuint16 != tuint162 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint16, tuint162)
	}

	var tInt32 int32 = -1000
	//int32ToBytes
	b = int32ToBytes(tInt32)
	tInt322 := bytesToInt32(b)
	if tInt32 != tInt322 {
		t.Fatalf("unexpected value obtained; got %b want %b", tInt32, tInt322)
	}
	//int32ToBytesBigEndian
	b = int32ToBytesBigEndian(tInt32)
	tInt322 = bytesToInt32BigEndian(b)
	if tInt32 != tInt322 {
		t.Fatalf("unexpected value obtained; got %b want %b", tInt32, tInt322)
	}

	//uint32ToBytes
	var tuint32 uint32 = 1000
	b = uint32ToBytes(tuint32)
	tuint322 := bytesToUint32(b)
	if tuint32 != tuint322 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint32, tuint322)
	}
	//uint32ToBytesBigEndian
	b = uint32ToBytesBigEndian(tuint32)
	tuint322 = bytesToUint32BigEndian(b)
	if tuint32 != tuint322 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint32, tuint322)
	}

	var tInt64 int64 = -1000
	//int64ToBytes
	b = int64ToBytes(tInt64)
	tInt642 := bytesToInt64(b)
	if tInt64 != tInt642 {
		t.Fatalf("unexpected value obtained; got %b want %b", tInt64, tInt642)
	}
	//int64ToBytesBigEndian
	b = int64ToBytesBigEndian(tInt64)
	tInt642 = bytesToInt64BigEndian(b)
	if tInt64 != tInt642 {
		t.Fatalf("unexpected value obtained; got %b want %b", tInt64, tInt642)
	}

	var tuint64 uint64 = 1000
	//uint64ToBytes
	b = uint64ToBytes(tuint64)
	tuint62 := bytesToUint64(b)
	if tuint64 != tuint62 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint64, tuint62)
	}
	//uint64ToBytesBigEndian
	b = uint64ToBytesBigEndian(tuint64)
	tuint62 = bytesToUint64BigEndian(b)
	if tuint64 != tuint62 {
		t.Fatalf("unexpected value obtained; got %b want %b", tuint64, tuint62)
	}

	var buf bytes.Buffer
	bb := NewBytesBuffer(buf)

	//WriteInt8 Readint8
	bb.Reset()
	bb.WriteInt8(-100)
	rd := NewFromBytes(bb.Bytes())
	yint8, _ := rd.ReadInt8()
	if yint8 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yint8, -100)
	}
	//WriteUint8 ReadUint8
	bb.Reset()
	bb.WriteUint8(100)
	rd = NewFromBytes(bb.Bytes())
	yuint8, _ := rd.ReadUint8()
	if yuint8 != 100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yuint8, 100)
	}

	//WriteInt16 ReadInt16
	bb.Reset()
	bb.WriteInt16(-100)
	rd = NewFromBytes(bb.Bytes())
	yint16, _ := rd.ReadInt16()
	if yint16 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yint16, -100)
	}

	//WriteInt16BigEndian ReadInt16BigEndian
	bb.Reset()
	bb.WriteInt16BigEndian(-100)
	rd = NewFromBytes(bb.Bytes())
	yint16, _ = rd.ReadInt16BigEndian()
	if yint16 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yint16, -100)
	}

	//WriteUint16 ReadUint16
	bb.Reset()
	bb.WriteUint16(100)
	rd = NewFromBytes(bb.Bytes())
	yuint16, _ := rd.ReadUint16()
	if yuint16 != 100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yuint16, 100)
	}

	//WriteUint16BigEndian ReadUint16
	bb.Reset()
	bb.WriteUint16BigEndian(100)
	rd = NewFromBytes(bb.Bytes())
	yuint16, _ = rd.ReadUint16BigEndian()
	if yuint16 != 100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yuint16, 100)
	}

	//WriteInt32 ReadInt32
	bb.Reset()
	bb.WriteInt32(-100)
	rd = NewFromBytes(bb.Bytes())
	yint32, _ := rd.ReadInt32()
	if yint32 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yint32, -100)
	}

	//WriteInt32BigEndian ReadInt32BigEndian
	bb.Reset()
	bb.WriteInt32BigEndian(-100)
	rd = NewFromBytes(bb.Bytes())
	yint32, _ = rd.ReadInt32BigEndian()
	if yint32 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yint32, -100)
	}

	//WriteInt32 ReadInt32
	bb.Reset()
	bb.WriteUint32(100)
	rd = NewFromBytes(bb.Bytes())
	yuint32, _ := rd.ReadUint32()
	if yint32 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yuint32, 100)
	}

	//WriteInt32 ReadInt32
	bb.Reset()
	bb.WriteUint32BigEndian(100)
	rd = NewFromBytes(bb.Bytes())
	yuint32, _ = rd.ReadUint32BigEndian()
	if yint32 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yuint32, 100)
	}

	//WriteInt64 ReadInt64
	bb.Reset()
	bb.WriteInt64(-100)
	rd = NewFromBytes(bb.Bytes())
	yint64, _ := rd.ReadInt64()
	if yint64 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yint64, -100)
	}

	//WriteInt64 ReadInt64
	bb.Reset()
	bb.WriteInt64BigEndian(-100)
	rd = NewFromBytes(bb.Bytes())
	yint64, _ = rd.ReadInt64BigEndian()
	if yint64 != -100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yint64, -100)
	}

	//WriteUint64 ReadUint64
	bb.Reset()
	bb.WriteUint64(100)
	rd = NewFromBytes(bb.Bytes())
	yuint64, _ := rd.ReadUint64()
	if yuint64 != 100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yuint64, 100)
	}

	//WriteUint64 ReadUint64
	bb.Reset()
	bb.WriteUint64BigEndian(100)
	rd = NewFromBytes(bb.Bytes())
	yuint64, _ = rd.ReadUint64BigEndian()
	if yuint64 != 100 {
		t.Fatalf("unexpected value obtained; got %v want %v", yuint64, 100)
	}

	//构造文本
	var s string
	for i := 0; i < 100; i++ {
		s = s + "s"
	}
	//WriteStringUint8 ReadStringUint8
	bb.Reset()
	bb.WriteStringUint8(s)
	rd = NewFromBytes(bb.Bytes())
	ss, _ := rd.ReadStringUint8()
	if ss != s {
		t.Fatalf("unexpected value obtained; got %v want %v", ss, s)
	}
	//WriteStringUint16 ReadStringUint16
	bb.Reset()
	bb.WriteStringUint16(s)
	rd = NewFromBytes(bb.Bytes())
	ss, _ = rd.ReadStringUint16()
	if ss != s {
		t.Fatalf("unexpected value obtained; got %v want %v", ss, s)
	}
	//WriteStringUint32 ReadStringUint32
	bb.Reset()
	bb.WriteStringUint32(s)
	rd = NewFromBytes(bb.Bytes())
	ss, _ = rd.ReadStringUint32()
	if ss != s {
		t.Fatalf("unexpected value obtained; got %v want %v", ss, s)
	}

	//WriteStringUint64 ReadStringUint64
	bb.Reset()
	bb.WriteStringUint64(s)
	rd = NewFromBytes(bb.Bytes())
	ss, _ = rd.ReadStringUint64()
	if ss != s {
		t.Fatalf("unexpected value obtained; got %v want %v", ss, s)
	}
}

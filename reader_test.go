package iox

import (
	"fmt"
	"io"
	"log"
	"testing"
)

func TestReaderSeeker(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.Ltime)
	buf := NewBytesBuffer()
	for i := 0; i < 255; i++ {
		buf.WriteInt8(int8(i))
	}
	rd := NewReadSeekerFromBytes(buf.Bytes())
	//CurPos
	curPos, err := rd.CurPos()
	if err != nil {
		t.Fatalf("unexpected value obtained; got %v want %v", err, nil)
	}
	if curPos != 0 {
		t.Fatalf("unexpected value obtained; got %v want %v", curPos, 0)
	}
	//Move
	err = rd.Move(100)
	if err != nil {
		t.Fatalf("unexpected value obtained; got %v want %v", err, nil)
	}
	curPos, err = rd.CurPos()
	if curPos != 100 {
		t.Fatalf("unexpected value obtained; got %b want %b", curPos, 0)
	}
	err = rd.Move(-101)
	if err == nil {
		t.Fatalf("unexpected value obtained; got %v want %v", err, fmt.Errorf("the legal pos range is between 0 and 254 ,and current pos is -1"))
	}
	//MoveTo
	err = rd.MoveTo(200)
	if err != nil {
		t.Fatalf("unexpected value obtained; got %v want %v", err, nil)
	}
	curPos, err = rd.CurPos()
	if curPos != 200 {
		t.Fatalf("unexpected value obtained; got %b want %b", curPos, 0)
	}
	//Len
	err = rd.MoveTo(0)
	if err != nil {
		panic(err)
	}
	length := rd.Size()
	if length != 255 {
		t.Fatalf("unexpected value obtained; got %b want %b", length, 255)
	}
	//LenUnRead
	length = rd.LenUnRead()
	if length != 255 {
		t.Fatalf("unexpected value obtained; got %b want %b", length, 255)
	}
	//readall
	for {
		_, err = rd.ReadBytes(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
	}
	length = rd.LenUnRead()
	if length != 0 {
		t.Fatalf("unexpected value obtained; got %b want %b", length, 0)
	}
	//ReadBytes
	rd.MoveTo(0)
	i := 0
	for {
		b, err := rd.ReadBytes(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if int(b[0]) != i {
			t.Fatalf("unexpected value obtained; got %v want %v", int(b[0]), i)
		}
		i++
	}
	//ReadBytesUnRead
	rd.Move(-2)
	b, err := rd.ReadBytesUnRead()
	if err != nil {
		panic(err)
	}
	if len(b) != 2 {
		t.Fatalf("unexpected value obtained; got %v want %v", len(b), 2)
	}
	if b[0] != 253 {
		t.Fatalf("unexpected value obtained; got %v want %v", int(b[0]), 253)
	}
	if b[1] != 254 {
		t.Fatalf("unexpected value obtained; got %v want %v", int(b[1]), 254)
	}
	wr := NewBytesBuffer()
	val := "test str"
	//ReadBytesUint8
	rd = nil
	wr.Reset()
	wr.WriteBytesUint8([]byte(val))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	val2, err := rd.ReadBytesUint8()
	if err != nil {
		panic(err)
	}
	if val != string(val2) {
		t.Fatalf("unexpected value obtained; got %v want %v", val, string(val2))
	}
	//ReadBytesUint16
	rd = nil
	wr.Reset()
	wr.WriteBytesUint16([]byte(val))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	val2, err = rd.ReadBytesUint16()
	if err != nil {
		panic(err)
	}
	if val != string(val2) {
		t.Fatalf("unexpected value obtained; got %v want %v", val, string(val2))
	}
	//ReadBytesUint16BigEndian
	rd = nil
	wr.Reset()
	wr.WriteBytesUint16BigEndian([]byte(val))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	val2, err = rd.ReadBytesUint16BigEndian()
	if err != nil {
		panic(err)
	}
	if val != string(val2) {
		t.Fatalf("unexpected value obtained; got %v want %v", val, string(val2))
	}
	//ReadBytesUint32
	rd = nil
	wr.Reset()
	wr.WriteBytesUint32([]byte(val))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	val2, err = rd.ReadBytesUint32()
	if err != nil {
		panic(err)
	}
	if val != string(val2) {
		t.Fatalf("unexpected value obtained; got %v want %v", val, string(val2))
	}
	//ReadBytesUint32BigEndian
	rd = nil
	wr.Reset()
	wr.WriteBytesUint32BigEndian([]byte(val))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	val2, err = rd.ReadBytesUint32BigEndian()
	if err != nil {
		panic(err)
	}
	if val != string(val2) {
		t.Fatalf("unexpected value obtained; got %v want %v", val, string(val2))
	}
	//WriteBytesUint64BigEndian
	rd = nil
	wr.Reset()
	wr.WriteBytesUint64([]byte(val))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	val2, err = rd.ReadBytesUint64()
	if err != nil {
		panic(err)
	}
	if val != string(val2) {
		t.Fatalf("unexpected value obtained; got %v want %v", val, string(val2))
	}
	//WriteBytesUint64BigEndian
	rd = nil
	wr.Reset()
	wr.WriteBytesUint64BigEndian([]byte(val))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	val2, err = rd.ReadBytesUint64BigEndian()
	if err != nil {
		panic(err)
	}
	if val != string(val2) {
		t.Fatalf("unexpected value obtained; got %v want %v", val, string(val2))
	}
	//ReadFloat32
	rd = nil
	wr.Reset()
	wr.WriteFloat32(float32(100.10))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	f32, err := rd.ReadFloat32()
	if err != nil {
		panic(err)
	}
	if f32 != float32(100.10) {
		t.Fatalf("unexpected value obtained; got %v want %v", f32, float32(100.10))
	}
	//ReadFloat32BigEndian
	rd = nil
	wr.Reset()
	wr.WriteFloat32BigEndian(float32(100.10))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	f32, err = rd.ReadFloat32BigEndian()
	if err != nil {
		panic(err)
	}
	if f32 != float32(100.10) {
		t.Fatalf("unexpected value obtained; got %v want %v", f32, float32(100.10))
	}
	//ReadFloat64
	rd = nil
	wr.Reset()
	wr.WriteFloat64(float64(100.10))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	f64, err := rd.ReadFloat64()
	if err != nil {
		panic(err)
	}
	if f64 != float64(100.10) {
		t.Fatalf("unexpected value obtained; got %v want %v", f64, float64(100.10))
	}

	//ReadFloat64
	rd = nil
	wr.Reset()
	wr.WriteFloat64BigEndian(float64(100.10))
	rd = NewReadSeekerFromBytes(wr.Bytes())
	f64, err = rd.ReadFloat64BigEndian()
	if err != nil {
		panic(err)
	}
	if f64 != float64(100.10) {
		t.Fatalf("unexpected value obtained; got %v want %v", f64, float64(100.10))
	}

	//create data
	rd = nil
	wr.Reset()
	sep := []byte{0x10}
	for i := 0; i < 4194304*10; i++ {
		if i%10000 == 0 {
			wr.WriteBytes(sep)
		} else {
			wr.WriteUint8(0)
		}
	}
	//LastIndexGen
	rd = NewReadSeekerFromBytes(wr.Bytes())
	index := rd.LastIndex(sep)
	if index != 41940000 {
		t.Fatalf("unexpected value obtained; got %v want %v", index, 41940000)
	}
	index = rd.IndexN(0, sep, 4195)
	if index != 41940000 {
		t.Fatalf("unexpected value obtained; got %v want %v", index, 41940000)
	}
	index = rd.IndexGen(1, 100, sep)
	if index != -1 {
		t.Fatalf("unexpected value obtained; got %v want %v", index, -1)
	}
	index = rd.LastIndexGen(1, 100, sep)
	if index != -1 {
		t.Fatalf("unexpected value obtained; got %v want %v", index, -1)
	}
	index = rd.LastIndexGen(0, 100, sep)
	if index != 0 {
		t.Fatalf("unexpected value obtained; got %v want %v", index, 0)
	}
	num := rd.Count(sep)
	if num != 4195 {
		t.Fatalf("unexpected value obtained; got %v want %v", num, 4195)
	}
}

package iox

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

// ReadSeeker helps you read and seek data from io.ReadSeeker,the default ByteOrder is LittleEndian.
type ReadSeeker struct {
	readSeeker io.ReadSeeker
}

//returns a *ReadSeeker from io.ReadSeeker.
func NewReadSeeker(rs io.ReadSeeker) *ReadSeeker {
	r := new(ReadSeeker)
	r.readSeeker = rs
	return r
}

//returns a *ReadSeeker from file.
func NewReadSeekerFromFile(fileName string) (*ReadSeeker, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0x666)
	if err != nil {
		return nil, err
	} else {
		return NewReadSeeker(file), nil
	}
}

//returns a *ReadSeeker from bytes.
func NewReadSeekerFromBytes(b []byte) *ReadSeeker {
	return NewReadSeeker(bytes.NewReader(b))
}

//Close if it's a io.Closer.
func (r *ReadSeeker) Close() {
	if closer, ok := r.readSeeker.(io.Closer); ok {
		closer.Close()
	}
}

//Seeking to an offset before the SeekCurrent of the file.
func (r *ReadSeeker) Move(n int64) error {
	curPos, err := r.readSeeker.Seek(0, io.SeekCurrent) //err always nil
	if err != nil {
		panic(err)
	}
	length, err := r.readSeeker.Seek(0, io.SeekEnd) //err always nil
	if err != nil {
		panic(err)
	}
	finalPos := curPos + n
	if !(finalPos >= 0 && finalPos < length) {
		_, err = r.readSeeker.Seek(curPos, io.SeekStart)
		if err != nil { //err always nil
			panic(err)
		}
		return fmt.Errorf("the legal pos range is between 0 and %v ,and current pos is %v", length-1, finalPos)
	}
	_, err = r.readSeeker.Seek(finalPos, io.SeekStart)
	if err != nil {
		return err
	}
	return nil
}

//Seeking to an offset before the start of the file.
func (r *ReadSeeker) MoveTo(pos int64) error {
	curPos, err := r.readSeeker.Seek(0, io.SeekCurrent) //err always nil
	if err != nil {
		panic(err)
	}
	length, err := r.readSeeker.Seek(0, io.SeekEnd) //err always nil
	if err != nil {
		panic(err)
	}
	if !(pos >= 0 && pos <= length) {
		_, err = r.readSeeker.Seek(curPos, io.SeekStart)
		if err != nil { //err always nil
			panic(err)
		}
		return fmt.Errorf("the legal pos range is between 0 and %v ,and current pos is %v", length-1, pos)
	}
	_, err = r.readSeeker.Seek(pos, io.SeekStart)
	if err != nil {
		return err
	}
	return nil
}

//get current pos of offset.
func (r *ReadSeeker) CurPos() (int64, error) {
	return r.readSeeker.Seek(0, io.SeekCurrent)
}

//get the size of the data.
func (r *ReadSeeker) Size() int64 {
	initialPos, err := r.CurPos()
	if err != nil {
		panic(err)
	}
	defer r.MoveTo(initialPos)
	n, err := r.readSeeker.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}
	return n
}

//get the length of unread data.
func (r *ReadSeeker) LenUnRead() int64 {
	curPos, err := r.CurPos()
	if err != nil {
		panic(err) //always nil
	}
	return r.Size() - curPos
}

//read n bytes.
func (r *ReadSeeker) ReadBytes(n int) ([]byte, error) {
	currentPos, err := r.CurPos()
	if err != nil {
		panic(err) //always nil
	}
	//check the surplus length of the data
	if surplusLen := r.Size() - currentPos; surplusLen < int64(n) {
		if surplusLen <= 0 {
			return nil, io.EOF
		} else {
			return nil, fmt.Errorf("%v is too long for this readSeeker,it's only %v bytes left,and the current position is:%v.", n, surplusLen, currentPos)
		}
	}
	bt := make([]byte, n)
	realRead, err := r.readSeeker.Read(bt)
	if err != nil {
		return nil, err
	}
	if realRead != n {
		return nil, fmt.Errorf("wish read not match real read")
	}
	return bt, nil
}

//read n bytes,just used in indexGen.
func (r *ReadSeeker) readBytesToDst(n int, dst []byte) ([]byte, error) {
	currentPos, err := r.CurPos()
	if err != nil {
		panic(err) //always nil
	}
	//check the surplus length of the data
	if surplusLen := r.Size() - currentPos; surplusLen < int64(n) {
		if surplusLen <= 0 {
			return nil, io.EOF
		} else {
			return nil, fmt.Errorf("%v is too long for this readSeeker,it's only %v bytes left,and the current position is:%v.", n, surplusLen, currentPos)
		}
	}
	realRead, err := r.readSeeker.Read(dst)
	if err != nil {
		return nil, err
	}
	if realRead != n {
		return nil, fmt.Errorf("wish read not match real read")
	}
	return dst, nil
}

//get all  unread data.
func (r *ReadSeeker) ReadBytesUnRead() ([]byte, error) {
	return r.ReadBytes(int(r.LenUnRead()))
}

//read uint8 as the data length and then read the data.
func (r *ReadSeeker) ReadBytesUint8() ([]byte, error) {
	n, err := r.ReadUint8()
	if err != nil {
		return nil, err
	}
	return r.ReadBytes(int(n))
}

//read uint16 as the data length and then read the data.
func (r *ReadSeeker) ReadBytesUint16() ([]byte, error) {
	n, err := r.ReadUint16()
	if err != nil {
		return nil, err
	}
	return r.ReadBytes(int(n))
}

//read uint16(BigEndian) as the data length and then read the data.
func (r *ReadSeeker) ReadBytesUint16BigEndian() ([]byte, error) {
	n, err := r.ReadUint16BigEndian()
	if err != nil {
		return nil, err
	}
	return r.ReadBytes(int(n))
}

//read uint32 as the data length and then read the data.
func (r *ReadSeeker) ReadBytesUint32() ([]byte, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	return r.ReadBytes(int(n))
}

//read uint32(BigEndian) as the data length and then read the data.
func (r *ReadSeeker) ReadBytesUint32BigEndian() ([]byte, error) {
	n, err := r.ReadUint32BigEndian()
	if err != nil {
		return nil, err
	}
	return r.ReadBytes(int(n))
}

//read uint64 as the data length and then read the data.
func (r *ReadSeeker) ReadBytesUint64() ([]byte, error) {
	n, err := r.ReadUint64()
	if err != nil {
		return nil, err
	}
	return r.ReadBytes(int(n))
}

//read uint64(BigEndian) as the data length and then read the data.
func (r *ReadSeeker) ReadBytesUint64BigEndian() ([]byte, error) {
	n, err := r.ReadUint64BigEndian()
	if err != nil {
		return nil, err
	}
	return r.ReadBytes(int(n))
}

//read uint8 as the data length and then read the data.
func (r *ReadSeeker) ReadStringUint8() (string, error) {
	n, err := r.ReadUint8()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//read uint16 as the data length and then read the data.
func (r *ReadSeeker) ReadStringUint16() (string, error) {
	n, err := r.ReadUint16()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//read uint16(BigEndian) as the data length and then read the data.
func (r *ReadSeeker) ReadStringUint16BigEndian() (string, error) {
	n, err := r.ReadUint16BigEndian()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//read uint32 as the data length and then read the data.
func (r *ReadSeeker) ReadStringUint32() (string, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//read uint32(BigEndian) as the data length and then read the data.
func (r *ReadSeeker) ReadStringUint32BigEndian() (string, error) {
	n, err := r.ReadUint32BigEndian()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//read uint64 as the data length and then read the data.
func (r *ReadSeeker) ReadStringUint64() (string, error) {
	n, err := r.ReadUint64()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//read uint64(BigEndian) as the data length and then read the data.
func (r *ReadSeeker) ReadStringUint64BigEndian() (string, error) {
	n, err := r.ReadUint64BigEndian()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//read n bytes of the data and then returns the hexadecimal encoding of the bytes.
func (r *ReadSeeker) ReadHexToString(n int) (string, error) {
	bt, err := r.ReadBytes(n)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString(bt)), nil
}

//read n bytes of the data and convert to string.
func (r *ReadSeeker) ReadString(n int) (string, error) {
	bt, err := r.ReadBytes(n)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

//read all unread data and convert to string.
func (r *ReadSeeker) ReadStringUnRead() (string, error) {
	bt, err := r.ReadBytesUnRead()
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

//read n bytes of data, then convert to string, and then remove spaces in the string.
func (r *ReadSeeker) ReadStringTrimSpace(n int) (string, error) {
	bt, err := r.ReadBytes(n)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bt)), nil

}

//read 1 byte and then convert to int8.
func (r *ReadSeeker) ReadInt8() (int8, error) {
	bt := make([]byte, 1)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	return int8(bt[0]), nil
}

//read 1 byte and then convert to uint8.
func (r *ReadSeeker) ReadUint8() (uint8, error) {
	bt := make([]byte, 1)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	return bt[0], nil
}

//read 2 bytes and then convert to int16.
func (r *ReadSeeker) ReadInt16() (int16, error) {
	bt := make([]byte, 2)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := int16(binary.LittleEndian.Uint16(bt))
	return n, nil
}

//read 2 bytes and then convert to int16(BigEndian).
func (r *ReadSeeker) ReadInt16BigEndian() (int16, error) {
	bt := make([]byte, 2)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := int16(binary.BigEndian.Uint16(bt))
	return n, nil
}

//read 2 bytes and then convert to uint16.
func (r *ReadSeeker) ReadUint16() (uint16, error) {
	bt := make([]byte, 2)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.LittleEndian.Uint16(bt)
	return n, nil
}

//read 2 bytes and then convert to uint16(BigEndian).
func (r *ReadSeeker) ReadUint16BigEndian() (uint16, error) {
	bt := make([]byte, 2)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.BigEndian.Uint16(bt)
	return n, nil
}

//read 4 bytes and then convert to int32.
func (r *ReadSeeker) ReadInt32() (int32, error) {
	bt := make([]byte, 4)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := int32(binary.LittleEndian.Uint32(bt))
	return n, nil
}

//read 4 bytes and then convert to int32(BigEndian).
func (r *ReadSeeker) ReadInt32BigEndian() (int32, error) {
	bt := make([]byte, 4)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := int32(binary.BigEndian.Uint32(bt))
	return n, nil
}

//read 4 bytes and then convert to uint32.
func (r *ReadSeeker) ReadUint32() (uint32, error) {
	bt := make([]byte, 4)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.LittleEndian.Uint32(bt)
	return n, nil
}

//read 4 bytes and then convert to uint32(BigEndian).
func (r *ReadSeeker) ReadUint32BigEndian() (uint32, error) {
	bt := make([]byte, 4)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.BigEndian.Uint32(bt)
	return n, nil
}

//read 8 bytes and then convert to int64.
func (r *ReadSeeker) ReadInt64() (int64, error) {
	bt := make([]byte, 8)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := int64(binary.LittleEndian.Uint64(bt))
	return n, nil
}

//read 8 bytes and then convert to int64(BigEndian).
func (r *ReadSeeker) ReadInt64BigEndian() (int64, error) {
	bt := make([]byte, 8)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := int64(binary.BigEndian.Uint64(bt))
	return n, nil
}

//read 8 bytes and then convert to uint64.
func (r *ReadSeeker) ReadUint64() (uint64, error) {
	bt := make([]byte, 8)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.LittleEndian.Uint64(bt)
	return n, nil
}

//read 8 bytes and then convert to uint64(BigEndian).
func (r *ReadSeeker) ReadUint64BigEndian() (uint64, error) {
	bt := make([]byte, 8)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.BigEndian.Uint64(bt)
	return n, nil
}

//read 4 bytes and convert it to float32.
func (r *ReadSeeker) ReadFloat32() (float32, error) {
	Uint32, err := r.ReadUint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(Uint32), nil
}

//read 4 bytes and convert it to float32(BigEndian).
func (r *ReadSeeker) ReadFloat32BigEndian() (float32, error) {
	Uint32, err := r.ReadUint32BigEndian()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(Uint32), nil
}

//read 8 bytes and convert it to float64.
func (r *ReadSeeker) ReadFloat64() (float64, error) {
	Uint64, err := r.ReadUint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(Uint64), nil
}

//read 8 bytes and convert it to float64(BigEndian).
func (r *ReadSeeker) ReadFloat64BigEndian() (float64, error) {
	Uint64, err := r.ReadUint64BigEndian()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(Uint64), nil
}

//Contains reports whether sep is within the data.
func (r *ReadSeeker) Contains(sep []byte) bool {
	return r.Index(sep) != -1
}

//Count counts the number of non-overlapping instances of sep in data.
func (r *ReadSeeker) Count(sep []byte) int64 {
	return r.CountGen(0, r.Size()-1, sep)
}

//Count counts the number of non-overlapping instances of sep in a range of data.
func (r *ReadSeeker) CountGen(beginPos, endPos int64, sep []byte) int64 {
	//防止文件位置发生移动
	initialPos, _ := r.CurPos()
	defer r.MoveTo(initialPos)
	lenSep := int64(len(sep))
	//sep输入不合法,标准库中是直接用f的长度加1，这里不照搬
	if lenSep == 0 {
		panic("sep can't be nil.")
	}
	if lenSep > r.Size() {
		return 0
	}
	curPos := beginPos
	var count, findPos int64
	r.MoveTo(curPos)
	for {
		if curPos > endPos {
			break
		}
		findPos = r.IndexGen(curPos, endPos, sep)
		if findPos == -1 {
			break
		} else {
			count++
			curPos = findPos + lenSep
		}
	}
	return count
}

//Index returns the index of the first instance of substr in data.
func (r *ReadSeeker) Index(sep []byte) int64 {
	endPos := r.Size() - 1
	return r.IndexGen(0, endPos, sep)
}

var bytesPoolForIndexGen = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 0, 1024)
		return buf
	},
}

//Index returns the index of the first instance of substr in a range of data.
func (r *ReadSeeker) IndexGen(beginPos, endPos int64, sep []byte) int64 {
	initialPos, _ := r.CurPos()
	defer r.MoveTo(initialPos)
	if realEndpos := r.Size() - 1; endPos < beginPos ||
		beginPos < 0 ||
		endPos < 0 ||
		realEndpos < endPos ||
		len(sep) == 0 {
		panic("beginPos:" + strconv.Itoa(int(beginPos)) + " or endPos:" + strconv.Itoa(int(endPos)) + " or sep is not a valid value.")
	}
	r.MoveTo(beginPos)
	nMaxSize := 1024
	lenSep := len(sep)
	if nMaxSize < lenSep*2 {
		nMaxSize = lenSep * 2 //the min size is lenSep * 2
	}
	buf := bytesPoolForIndexGen.Get().([]byte)
	defer bytesPoolForIndexGen.Put(buf)
	for {
		buf = buf[0:0]
		curPos, _ := r.CurPos()
		if int(endPos-curPos+1) < nMaxSize {
			buf = make([]byte, int(endPos-curPos+1))
			buf, _ = r.readBytesToDst(int(endPos-curPos+1), buf)
		} else {
			buf = make([]byte, nMaxSize)
			buf, _ = r.readBytesToDst(nMaxSize, buf)
		}
		newPos := bytes.Index(buf, sep)
		if len(buf) < nMaxSize {
			if newPos == -1 {
				return -1
			} else {
				return curPos + int64(newPos)
			}
		} else {
			if newPos >= 0 {
				return curPos + int64(newPos)
			}
		}
		r.Move(int64(0 - lenSep))
	}
	return -1
}

//Index returns the nth index of the instance of sep in data.
func (r *ReadSeeker) IndexN(beginPos int64, sep []byte, n int) int64 {
	initialPos, _ := r.CurPos()
	defer r.MoveTo(initialPos)
	r.MoveTo(beginPos)
	if n <= 0 {
		panic(strconv.Itoa(n) + " is not a valid value.")
	}
	var findPos int64
	endPos := r.Size() - 1
	for i := 0; i < n; i++ {
		curPos, _ := r.CurPos()
		if curPos > endPos {
			return -1
		}
		findPos = r.IndexGen(curPos, endPos, sep)
		if findPos == -1 {
			return -1
		} else {
			r.MoveTo(findPos + int64(len(sep)))
		}
	}
	return findPos
}

//LastIndex returns the index of the last instance of sep in data.
func (r *ReadSeeker) LastIndex(sep []byte) int64 {
	endPos := r.Size() - 1
	return r.LastIndexGen(0, endPos, sep)
}

//LastIndex returns the index of the last instance of sep in a range of data.
func (r *ReadSeeker) LastIndexGen(beginPos, endPos int64, sep []byte) int64 {
	initialPos, _ := r.CurPos()
	defer r.MoveTo(initialPos)
	if realEndpos := r.Size() - 1; endPos < beginPos ||
		beginPos < 0 ||
		endPos < 0 ||
		realEndpos < endPos ||
		len(sep) == 0 {
		panic("beginPos:" + strconv.Itoa(int(beginPos)) + " or endPos:" + strconv.Itoa(int(endPos)) + " or sep is not a valid value.")
	}
	r.MoveTo(endPos + 1)
	nMaxSize := 1024
	lenSep := len(sep)
	if nMaxSize < lenSep*2 {
		nMaxSize = lenSep * 2
	}
	buf := bytesPoolForIndexGen.Get().([]byte)
	defer bytesPoolForIndexGen.Put(buf)
	for {
		buf = buf[:0]
		curPos, _ := r.CurPos()
		if int(curPos-beginPos) < nMaxSize {
			buf = make([]byte, int(curPos-beginPos))
			buf, _ = r.ReadBytesReverse(int(curPos - beginPos))
		} else {
			buf = make([]byte, nMaxSize)
			buf, _ = r.ReadBytesReverse(nMaxSize)
		}
		newPos := bytes.LastIndex(buf, sep)
		tempPos, _ := r.CurPos()
		if len(buf) < nMaxSize {
			if newPos == -1 {
				return -1
			} else {
				return beginPos + int64(newPos)
			}
		} else {
			if newPos >= 0 {
				return tempPos + int64(newPos)
			}
		}
		r.Move(int64(lenSep))
	}
	return -1
}

//read n bytes,read the data backwards.
func (r *ReadSeeker) ReadBytesReverse(n int) ([]byte, error) {
	err := r.Move(int64(-n))
	if err != nil {
		return nil, err
	}
	currentPos, err := r.CurPos()
	if err != nil {
		return nil, err
	}
	if surplusLen := r.Size() - currentPos; surplusLen < int64(n) {
		if surplusLen < 0 {
			surplusLen = 0
		}
		if surplusLen == 0 {
			return nil, io.EOF
		} else {
			return nil, fmt.Errorf(fmt.Sprint(n, ` is too long for this readSeeker,it's only `, surplusLen, ` byte left,and the current position is:`, currentPos, `.`))
		}
	}
	defer r.MoveTo(currentPos)
	bt := make([]byte, n)
	realRead, err := r.readSeeker.Read(bt)
	if err != nil {
		return nil, err
	}
	if realRead != n {
		return nil, fmt.Errorf("wish read not match real read")
	}
	return bt, nil
}

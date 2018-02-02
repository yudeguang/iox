//主要用于在文件(os.File)中查找相关字节数组 Index(b []byte) 以及读取数据 ReadByte(n int)等操作，同时也支持所有io.ReadSeeker类型数据
package files

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//ReadSeeker及为io.ReadSeeker
type ReadSeeker struct {
	readSeeker io.ReadSeeker
}

//从io.ReadSeeker初始化
func New(ioReadSeeker io.ReadSeeker) *ReadSeeker {
	r := new(ReadSeeker)
	r.readSeeker = ioReadSeeker
	return r
}

//从文件初始化,注意最后要调用Close关闭文件
func NewFromFile(fileName string) (*ReadSeeker, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0x666)
	if err != nil {
		return nil, err
	} else {
		return New(file), nil
	}
}

//从切片初始化
func NewFromBytes(b []byte) *ReadSeeker {
	return New(bytes.NewReader(b))
}

//关闭f
func (r *ReadSeeker) Close() {
	if closer, ok := r.readSeeker.(io.Closer); ok {
		closer.Close()
	}
}

//从当前位置移动多少个字节
func (r *ReadSeeker) Move(n int64) error {
	_, err := r.readSeeker.Seek(n, io.SeekCurrent)
	if err != nil {
		return err
	}
	return nil
}

//移动到某个位置
func (r *ReadSeeker) MoveTo(pos int64) error {
	_, err := r.readSeeker.Seek(pos, io.SeekStart)
	if err != nil {
		return err
	}
	return nil
}

//获得当前位置
func (r *ReadSeeker) CurPos() (int64, error) {
	return r.readSeeker.Seek(0, io.SeekCurrent)
}

//获得长度
func (r *ReadSeeker) Size() int64 {
	//防止文件位置发生移动
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

//获得长度Size()函数的别名
func (r *ReadSeeker) Len() int64 {
	return r.Size()
}

//获取剩余未读取的长度
func (r *ReadSeeker) LenUnRead() int64 {
	curPos, err := r.CurPos()
	if err != nil {
		panic(err)
	}
	return r.Size() - curPos
}

//读取n个字节,并移动指针
func (r *ReadSeeker) ReadByte(n int) ([]byte, error) {
	currentPos, err := r.CurPos()
	if err != nil {
		return nil, err
	}
	//强制检查传入参数的合法性
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

//从当前位置起，读取剩余的
func (r *ReadSeeker) ReadByteUnRead() ([]byte, error) {
	return r.ReadByte(int(r.LenUnRead()))
}

//读取若干字节,并移动指针,其中第1个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadByteUint8() ([]byte, error) {
	n, err := r.ReadUint8()
	if err != nil {
		return nil, err
	}
	return r.ReadByte(int(n))
}

//读取若干字节,并移动指针,其中前2个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadByteUint16() ([]byte, error) {
	n, err := r.ReadUint16()
	if err != nil {
		return nil, err
	}
	return r.ReadByte(int(n))
}

//读取若干字节,并移动指针,其中前2个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadByteUint16BigEndian() ([]byte, error) {
	n, err := r.ReadUint16BigEndian()
	if err != nil {
		return nil, err
	}
	return r.ReadByte(int(n))
}

//读取若干字节,并移动指针,其中前4个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadByteUint32() ([]byte, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return nil, err
	}
	return r.ReadByte(int(n))
}

//读取若干字节,并移动指针,其中前4个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadByteUint32BigEndian() ([]byte, error) {
	n, err := r.ReadUint32BigEndian()
	if err != nil {
		return nil, err
	}
	return r.ReadByte(int(n))
}

//读取若干字节,并移动指针,其中前8个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadByteUint64() ([]byte, error) {
	n, err := r.ReadUint64()
	if err != nil {
		return nil, err
	}
	return r.ReadByte(int(n))
}

//读取若干字节,并移动指针,其中前8个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadByteUint64BigEndian() ([]byte, error) {
	n, err := r.ReadUint64BigEndian()
	if err != nil {
		return nil, err
	}
	return r.ReadByte(int(n))
}

//读取若干字节,转化为文本,并移动指针,其中第1个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadStringUint8() (string, error) {
	n, err := r.ReadUint8()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//读取若干字节,转化为文本,并移动指针,其中前2个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadStringUint16() (string, error) {
	n, err := r.ReadUint16()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//读取若干字节,转化为文本,并移动指针,其中前2个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadStringUint16BigEndian() (string, error) {
	n, err := r.ReadUint16BigEndian()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//读取若干字节,转化为文本,并移动指针,其中前4个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadStringUint32() (string, error) {
	n, err := r.ReadUint32()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//读取若干字节,转化为文本,并移动指针,其中前4个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadStringUint32BigEndian() (string, error) {
	n, err := r.ReadUint32BigEndian()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//读取若干字节,转化为文本,并移动指针,其中前8个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadStringUint64() (string, error) {
	n, err := r.ReadUint64()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//读取若干字节,转化为文本,并移动指针,其中前8个字节表示后续待读取切片长度
func (r *ReadSeeker) ReadStringUint64BigEndian() (string, error) {
	n, err := r.ReadUint64BigEndian()
	if err != nil {
		return "", err
	}
	return r.ReadString(int(n))
}

//读取n个字节的16进制数据,转化为string,并移动指针
func (r *ReadSeeker) ReadHexToString(n int) (string, error) {
	bt, err := r.ReadByte(n)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(hex.EncodeToString((bt))), nil
}

//读取n个字节,转化为string,并移动指针
func (r *ReadSeeker) ReadString(n int) (string, error) {
	bt, err := r.ReadByte(n)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

//从当前位置开始，读取剩余部分数据
func (r *ReadSeeker) ReadStringUnRead() (string, error) {
	bt, err := r.ReadByteUnRead()
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

//读取n个字节,转化为去前后空格的string,并移动指针
func (r *ReadSeeker) ReadStringTrimSpace(n int) (string, error) {
	bt, err := r.ReadByte(n)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bt)), nil

}

//读取一个uint8数,并移动指针
func (r *ReadSeeker) ReadUint8() (uint8, error) {
	bt := make([]byte, 1)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	return bt[0], nil
}

//读取一个uint16数,并移动指针
func (r *ReadSeeker) ReadUint16() (uint16, error) {
	bt := make([]byte, 2)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.LittleEndian.Uint16(bt)
	return n, nil
}

//读取一个uint16数,并移动指针
func (r *ReadSeeker) ReadUint16BigEndian() (uint16, error) {
	bt := make([]byte, 2)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.BigEndian.Uint16(bt)
	return n, nil
}

//读取一个uint32数,并移动指针
func (r *ReadSeeker) ReadUint32() (uint32, error) {
	bt := make([]byte, 4)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.LittleEndian.Uint32(bt)
	return n, nil
}

//读取一个uint32数,并移动指针
func (r *ReadSeeker) ReadUint32BigEndian() (uint32, error) {
	bt := make([]byte, 4)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.BigEndian.Uint32(bt)
	return n, nil
}

//读取一个uint64数,并移动指针
func (r *ReadSeeker) ReadUint64() (uint64, error) {
	bt := make([]byte, 8)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.LittleEndian.Uint64(bt)
	return n, nil
}

//读取一个uint64数,并移动指针
func (r *ReadSeeker) ReadUint64BigEndian() (uint64, error) {
	bt := make([]byte, 8)
	_, err := r.readSeeker.Read(bt)
	if err != nil {
		return 0, err
	}
	n := binary.BigEndian.Uint64(bt)
	return n, nil
}

//判断f是否包含子串sep
func (r *ReadSeeker) Contains(sep []byte) bool {
	return r.Index(sep) != -1
}

//计算f中共有多少个不重重叠的sep子串
func (r *ReadSeeker) Count(sep []byte) int64 {
	return r.CountGen(0, r.Size()-1, sep)
}

//判断f中某段数据内有多少个不重叠的sep子串
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

//查找数据,从文件的开始位置开始查找sep子串,返回第1次出现的位置
func (r *ReadSeeker) Index(sep []byte) int64 {
	endPos := r.Size() - 1
	return r.IndexGen(0, endPos, sep)
}

//查找数据,从文件的beginPos位置到endPos位置开始查找sep子串,返回第1次出现的位置
//注意待搜索数据包含beginPos及endPos本身
func (r *ReadSeeker) IndexGen(beginPos, endPos int64, sep []byte) int64 {
	//防止文件位置发生移动
	initialPos, _ := r.CurPos()
	defer r.MoveTo(initialPos)

	//先对输入值合法性进行判断,不合法则直接退出,强制用户检察
	if realEndpos := r.Size() - 1; endPos < beginPos ||
		beginPos < 0 ||
		endPos < 0 ||
		realEndpos < endPos ||
		len(sep) == 0 {
		panic("beginPos:" + strconv.Itoa(int(beginPos)) + " or endPos:" + strconv.Itoa(int(endPos)) + " or sep is not a valid value.")
		//return -1
	}
	r.MoveTo(beginPos)
	//每次实际读取的字节数组长度,至少要是2倍于待对比的Sep长度
	//除非剩余数据不够长，那么默认每次读取的最小磁盘大小为4M(4 * 1024 * 1024)
	nMaxSize := 4194304
	lenSep := len(sep)
	if nMaxSize < lenSep*2 {
		nMaxSize = lenSep * 2
	}
	for {
		var data []byte
		curPos, _ := r.CurPos()
		//读取每次待对比的数据,如果剩余数据太短,则读取到endPos为止
		if int(endPos-curPos+1) < nMaxSize {
			data, _ = r.ReadByte(int(endPos - curPos + 1))
		} else {
			data, _ = r.ReadByte(nMaxSize)
		}
		newPos := bytes.Index(data, sep)
		if len(data) < nMaxSize { //末次读取
			if newPos == -1 {
				return -1
			} else {
				return curPos + int64(newPos)
			}
		} else { //非末次读取
			if newPos >= 0 {
				return curPos + int64(newPos)
			}
		}
		//文件位置前移lenSep个长度
		r.Move(int64(0 - lenSep))
	}
	return -1
}

//查找数据,从文件的beginPos位置开始查找S子串,第N次出现的位置,N大于0,beginPos大于等于0
func (r *ReadSeeker) IndexN(beginPos int64, sep []byte, n int) int64 {
	//防止文件位置发生移动
	initialPos, _ := r.CurPos()
	defer r.MoveTo(initialPos)
	r.MoveTo(beginPos)
	// 检察N的合法性
	if n <= 0 {
		panic(strconv.Itoa(n) + " is not a valid value.")
		// return -1
	}
	// 执行n次搜索
	var findPos int64
	endPos := r.Size() - 1
	for i := 0; i < n; i++ {
		curPos, _ := r.CurPos()
		if curPos > endPos {
			return -1
		}
		findPos = r.IndexGen(curPos, endPos, sep)
		if findPos == -1 { //中间任何一次没找到就返回
			return -1
		} else { //找到一次,那么要移动到这个位置加SEP长度
			r.MoveTo(findPos + int64(len(sep)))
		}
	}
	return findPos
}

//查找数据,在整个文件中查找,返回最后一次出现sep的位置,注意待搜索数据包含beginPos及endPos本身
func (r *ReadSeeker) LastIndex(sep []byte) int64 {
	endPos := r.Size() - 1
	return r.LastIndexGen(0, endPos, sep)
}

//查找数据,在的beginPos到endPos内查找sep子串,返回最后1次出现的位置,注意待搜索数据包含beginPos及endPos本身
func (r *ReadSeeker) LastIndexGen(beginPos, endPos int64, sep []byte) int64 {
	//防止文件位置发生移动
	initialPos, _ := r.CurPos()
	defer r.MoveTo(initialPos)
	//先对输入值合法性进行判断,不合法则直接退出,强制用户检察
	if realEndpos := r.Size() - 1; endPos < beginPos ||
		beginPos < 0 ||
		endPos < 0 ||
		realEndpos < endPos ||
		len(sep) == 0 {
		panic("beginPos:" + strconv.Itoa(int(beginPos)) + " or endPos:" + strconv.Itoa(int(endPos)) + " or sep is not a valid value.")
		//return -1
	}
	//这种方法是移动到最后
	r.MoveTo(endPos + 1)
	//每次实际读取的字节数组长度,至少要是2倍于待对比的Sep长度
	//除非剩余数据不够长，那么默认每次读取的最小磁盘大小为4M(4 * 1024 * 1024)
	nMaxSize := 4194304
	lenSep := len(sep)
	if nMaxSize < lenSep*2 {
		nMaxSize = lenSep * 2
	}
	for {
		var data []byte
		curPos, _ := r.CurPos()
		//读取每次待对比的数据,如果剩余数据太短,则读取到beginPos为止
		if int(curPos-beginPos) < nMaxSize {
			r.MoveTo(beginPos)
			data, _ = r.ReadByte(int(curPos - beginPos))
		} else {
			data, _ = r.ReadByteReverse(nMaxSize)
		}
		newPos := bytes.LastIndex(data, sep)
		tempPos, _ := r.CurPos()
		if len(data) < nMaxSize { //末次读取
			if newPos == -1 {
				return -1
			} else {
				return beginPos + int64(newPos)
			}
		} else { //非末次读取
			if newPos >= 0 {
				return tempPos + int64(newPos)
			}
		}
		//文件位置往后移lenSep个长度
		r.Move(int64(lenSep))
	}
	return -1
}

//从当前位置向前读取n个字节,并移动指针
func (r *ReadSeeker) ReadByteReverse(n int) ([]byte, error) {
	//先向前移动n个位置
	err := r.Move(int64(-n))
	if err != nil {
		return nil, err
	}
	//最终读取后指针所在位置
	currentPos, err := r.CurPos()
	if err != nil {
		return nil, err
	}
	//强制检查传入参数的合法性
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

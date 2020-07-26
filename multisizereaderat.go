package ko

import (
	"io"
)

type SizeReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
	Size() int64
}

type MultiSizeReaderAt struct {
	readers []SizeReaderAt
	size  int64
}

func NewMultiSizeReaderAt(rs ...SizeReaderAt) SizeReaderAt {
	var size int64
	for _, v := range rs {
		size += v.Size()
	}
	return &MultiSizeReaderAt{
		readers: rs,
		size:    size,
	}
}

func (m MultiSizeReaderAt) Size() int64{
	return m.size
}

func (m *MultiSizeReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	full := len(p)
	if off >= m.size {
		return 0, io.EOF
	}
	for i:=0; i < len(m.readers) && len(p) !=0; i++{
		if int(m.readers[i].Size()) < int(off){
			off -= m.readers[i].Size()
			continue
		}
		k, err :=m.readers[i].ReadAt(p, off)
		n +=  k
		if err != nil {
			if err.Error() != "EOF" {
				return n, err
			}
			if i == len(m.readers)-1 && err.Error() == "EOF" {
				return n, io.EOF
			}
		}
		p = p[k:]
		off = 0
	}
	if n != full {
		return n, io.EOF
	}
	return n, nil
}
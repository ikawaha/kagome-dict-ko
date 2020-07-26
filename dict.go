package ko

import (
	"archive/zip"
	"bytes"
	"fmt"
	"sort"
	"sync"

	"github.com/ikawaha/kagome-dict-ko/internal/data"
	"github.com/ikawaha/kagome/v2/dict"
)

type systemDict struct {
	once sync.Once
	dict *dict.Dict
}

var (
	full   systemDict
	shrink systemDict
)

// Dict returns a dictionary.
func Dict() *dict.Dict {
	full.once.Do(func() {
		full.dict = loadDict(true)
		shrink.once.Do(func() {
			shrink.dict = full.dict
		})
	})
	return full.dict
}

// DictShrink returns a dictionary without content part.
// note. If an unshrinked dictionary already exists, this function returns it.
func DictShrink() *dict.Dict {
	shrink.once.Do(func() {
		shrink.dict = loadDict(false)
	})
	return shrink.dict
}

func loadDict(full bool) (d *dict.Dict) {
	pieces := data.AssetNames()
	sort.Strings(pieces)

	rs := make([]SizeReaderAt, 0, len(pieces))
	for _, v := range pieces {
		b, err := data.Asset(v)
		if err != nil {
			panic(fmt.Errorf("assert error, %q, %v", v, err))
		}
		rs = append(rs, bytes.NewReader(b))
	}
	r := NewMultiSizeReaderAt(rs...)
	zr, err := zip.NewReader(r, r.Size())
	if err != nil {
		panic(err)
	}
	d, err = dict.Load(zr, full)
	if err != nil {
		panic(err)
	}
	return d
}

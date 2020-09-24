package ko

import (
	"archive/zip"
	"bytes"
	"fmt"
	"sort"
	"sync"

	"github.com/ikawaha/kagome-dict-ko/internal/data"
	"github.com/ikawaha/kagome-dict/dict"
)

type FeatureIndex int

const (
	// Features are information given to a word, such as follows:
	// 공원	NNG,장소,T,공원,*,*,*,*
	// 에	JKB,*,F,에,*,*,*,*
	// 갔	VV+EP,*,T,갔,Inflect,VV,EP,가/VV/*+았/EP/*
	// 다	EF,*,F,다,*,*,*,*
	// .	SF,*,*,*,*,*,*,*
	// EOS
	// POSHierarchy represents part-of-speech hierarchy
	// e.g. Columns NNG POSs which hierarchy depth is 1.
	POSHierarchy = 1
	// Meaning
	Meaning FeatureIndex = 1
	// Presence or absence, T for true; F for false; else *.
	PresenceOrAbsence = 2
	// Reading, usually matches surface, but may differ for foreign words e.g. Chinese character words.
	Reading = 3
	// Type, one of: Inflect (활용); Compound (복합명사); or Preanalysis (기분석).
	Type = 4
	// FirstPOS, e.g. given a part-of-speech tag of "VV+EM+VX+EP", would return VV.
	FirstPOS = 5
	// LastPOS, e.g. given a part-of-speech tag of "VV+EM+VX+EP", would return EP.
	LastPOS = 6
	// Expression, fields that tell how usage, compound nouns, and key analysis are organized.
	Expression = 7
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

	rs := make([]dict.SizeReaderAt, 0, len(pieces))
	for _, v := range pieces {
		b, err := data.Asset(v)
		if err != nil {
			panic(fmt.Errorf("assert error, %q, %v", v, err))
		}
		rs = append(rs, bytes.NewReader(b))
	}
	r := dict.MultiSizeReaderAt(rs...)
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

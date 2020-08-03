package ko

import (
	"testing"

	"github.com/ikawaha/kagome/v2/dict"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

const (
	KoDictEntrySize = 816283
	testDictPath    = "testdata/ko.dict"
)

func TestDictShrink(t *testing.T) {
	d := DictShrink()
	if want, got := KoDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := 0, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLoadShrink(t *testing.T) {
	d, err := dict.LoadShrink(testDictPath)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := KoDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := 0, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestNew(t *testing.T) {
	d := Dict()
	if want, got := KoDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := KoDictEntrySize, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLoadDictFile(t *testing.T) {
	d, err := dict.LoadDictFile(testDictPath)
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	if want, got := KoDictEntrySize, len(d.Morphs); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
	if want, got := KoDictEntrySize, len(d.Contents); want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSingleton(t *testing.T) {
	a := Dict()
	b := Dict()
	if a != b {
		t.Errorf("got %p and %p, expected singleton", a, b)
	}
}

func TestContentsMeta(t *testing.T) {
	d := Dict()
	if want, got := 0, d.ContentsMeta[dict.POSStartIndex]; want != int(got) {
		t.Errorf("pos start index: want %d, got %d", want, got)
	}
	if want, got := 1, d.ContentsMeta[dict.POSEndIndex]; want != int(got) {
		t.Errorf("pos end index: want %d, got %d", want, got)
	}
	if v := d.ContentsMeta[dict.BaseFormIndex]; v >= 0 {
		t.Errorf("base form index: want <0, got %v", v)
	}
	if want, got := 3, d.ContentsMeta[dict.ReadingIndex]; want != int(got) {
		t.Errorf("reading index: want %d, got %d", want, got)
	}
	if v := d.ContentsMeta[dict.PronunciationIndex]; v >= 0 {
		t.Errorf("pronunciation index: want <0, got %v", v)
	}
}

func TestReading(t *testing.T) {
	tnz, err := tokenizer.New(Dict())
	if err != nil {
		t.Fatalf("unexpected error, %v", err)
	}
	tokens := tnz.Tokenize("초밥이먹고싶다.")
	if want, got := 9, len(tokens); want != got {
		t.Fatalf("token length: want %d, got %d", want, got)
	}
	//BOS
	if got, ok := tokens[0].Reading(); ok {
		t.Errorf("want !ok, got %q", got)
	}
	//초밥->초밥
	if got, ok := tokens[1].Reading(); !ok {
		t.Error("want ok, but !ok")
	} else if want := "초밥"; want != got {
		t.Fatalf("want %s, got %s", want, got)
	}
}

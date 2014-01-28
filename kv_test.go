package log

import (
	"testing"
	"time"
)

type foo struct {
	I   int
	S   string
	R   rune
	B   bool
	T   time.Time
	D   time.Duration
	foo bool // unexported, so won't be logged
	Foo bool `kv:"-"`      // "-" means that the field won't be logged
	Bar bool `kv:"foobar"` // change the key
}

type bar struct {
	Bar int
	Foo *foo
}

var f = &foo{
	I: 34,
	S: "all aboard!",
	R: 'b',
	T: time.Date(2006, 8, 24, 02, 30, 0, 0, time.UTC),
	D: 30 * time.Millisecond}

var b = &bar{Bar: 0, Foo: f}

func TestMarshalStruct(t *testing.T) {
	m, err := Marshal(f)
	if err != nil {
		t.Error("marshal failed", err)
	}
	ex := "I=34\tS=\"all aboard!\"\tR='b'\tB=false\tD=30ms\tfoobar=false"
	if string(m) != ex {
		t.Error("Marhshal returned a wrong result")
	}

	m, err = Marshal(*f)
	if err != nil {
		t.Error("marshal failed", err)
	}
	if string(m) != ex {
		t.Error("Marhshal returned a wrong result")
	}
}

func TestMarshalDeepStruct(t *testing.T) {
	m, err := Marshal(b)
	if err != nil {
		t.Error("marshal failed", err)
	}
	ex := "Bar=0\tFoo.I=34\tFoo.S=\"all aboard!\"\tFoo.R='b'\tFoo.B=false\tFoo.D=30ms\tFoo.foobar=false"
	if string(m) != ex {
		t.Error("Marhshal returned a wrong result")
	}
}

func TestMarshalMap(t *testing.T) {
	m, err := Marshal(map[string]interface{}{
		"foo":  true,
		"rune": 'c',
		"bar":  time.Date(2006, 8, 24, 02, 30, 0, 0, time.UTC),
	})
	if err != nil {
		t.Error("marshal failed.", err)
	}
	if string(m) != "foo=true\trune='c'\tbar=2006-08-24T02:30:00Z" {
		t.Error("marshalisation is wrong")
	}
}

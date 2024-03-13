package safeconv

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestConvert(t *testing.T) {
	for _, test := range []convertOKTestCase[int, float64]{
		{v: 1, want: 1, wantOk: true},
		{v: 2 << 53, want: 2 << 53, wantOk: true},
		{v: -2 << 53, want: -2 << 53, wantOk: true},
		{v: 2<<53 + 2, wantOk: false},
		{v: -2<<53 - 2, wantOk: false},
		{v: 2<<53 + 4, want: 2<<53 + 4, wantOk: true},
	} {
		assertConvert(t, test)
	}

	for _, test := range []convertOKTestCase[float32, float64]{
		{v: 1, want: 1, wantOk: true},
		{v: 0.1, want: 0.10000000149011612, wantOk: true}, // TODO: Do we want to allow this?
		{v: float32(math.NaN()), wantOk: false},
	} {
		assertConvert(t, test)
	}

	for _, test := range []convertOKTestCase[float64, float32]{
		{v: 1, want: 1, wantOk: true},
		{v: 0.1, wantOk: false},
		{v: math.NaN(), wantOk: false},
	} {
		assertConvert(t, test)
	}

	for _, test := range []convertOKTestCase[float64, float32]{
		{v: 1, want: 1, wantOk: true},
		{v: 0.1, wantOk: false},
		{v: math.NaN(), wantOk: false},
	} {
		assertConvert(t, test)
	}

	for _, test := range []convertOKTestCase[float64, time.Duration]{
		{v: 1, want: 1, wantOk: true},
		{v: 0.1, wantOk: false},
		{v: math.NaN(), wantOk: false},
	} {
		assertConvert(t, test)
	}
}

type convertOKTestCase[T, U number] struct {
	v      T
	want   U
	wantOk bool
}

func assertConvert[T, U number](t *testing.T, test convertOKTestCase[T, U]) {
	desc := fmt.Sprintf("Convert(%[1]T(%[1]v)) = %[2]v, %[3]v", test.v, test.want, test.wantOk)
	t.Run(desc, func(t *testing.T) {
		got, gotOK := ConvertOK[T, U](test.v)
		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
		if gotOK != test.wantOk {
			t.Errorf("got OK %v, want OK %v", gotOK, test.wantOk)
		}
	})
}

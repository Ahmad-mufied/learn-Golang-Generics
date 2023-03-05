package learn_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
==	Type Approximation

--	Kadang, kita sering membuat Type Declaration di Golang untuk tipe data lain, misal kita membuat tipe data Age untuk tipe data int
--	Secara default, jika kita gunakan Age sebagai type declaration untuk int, lalu kita membuat Type Set yang berisi constraint int, maka tipe data Age dianggap tidak compatible dengan Type Set yang kita buat

*/

type Age int

type Number1 interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min1[T Number1](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin1(t *testing.T) {
	assert.Equal(t, 100, Min1[int](100, 200))
	assert.Equal(t, int64(100), Min1[int64](int64(100), int64(200)))
	assert.Equal(t, float64(100), Min1[float64](float64(100), float64(200)))
	assert.Equal(t, Age(100), Min1[Age](Age(100), Age(200)))
}

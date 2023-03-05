package learn_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
==	In Line Type Constraint

--	Sebelum-sebelumnya, kita selalu menggunakan type declaration atau type set ketika membuat type constraint di type parameter
--	Sebenarnya tidak ada kewajiban kita harus membuat type declaration atau type set jika kita ingin membuat type parameter,
	kita bisa gunakan secara langsung (in line) pada type constraint, misalnya di awal kita sudah bahas tentang interface {} (kosong),
	tapi kita selalu gunakan type declaration any
--	Jika kita mau, kita juga bisa langsung gunakan interface { int | float32 | float64} dibanding membuat type set Number misalnya

*/

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 200))
	assert.Equal(t, int64(100), FindMin[int64](int64(100), int64(200)))
}

/*
==	Generic di Type Parameter

--	Pada kasus tertentu, kadang ada kebutuhan kita menggunakan type parameter yang ternyata type ternyata type tersebut juga generic atau memiliki type parameter
--	Kita juga bisa menggunakan in line type constraint agar lebih mudah, dengan cara menambahkan type parameter selanjutnya, misal
--	[S interface{[]E}, E interface{}], artinya S harus slice element E, dimana E boleh tipe apapun
--	[S []E, E any], artinya S harus slice element E, dimana E boleh tipe apapun

*/

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Ahmad", "Mufied", "Nugroho",
	}

	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Ahmad", first)
}

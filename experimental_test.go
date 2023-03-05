package learn_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"testing"
)

/*
==	Experimental Package

--	Saat versi Go-Lang 1.18, terdapat experimental package yang banyak menggunakan fitur Generic, namun belum resmi masuk ke Go-Lang Standard Library
--	Kedepannya, karena ini masih experimental (percobaan), bisa jadi package ini akan berubah atau bahkan mungkin akan dihapus
--	https://pkg.go.dev/golang.org/x/exp
--	Silahkan install sebagai dependency di Go Modules menggunakan perintah
--	go get golang.org/x/exp

==	Constraints Package

--	Constraints Package berisi type declaration yang bisa kita gunakan untuk tipe data bawaan Go-Lang, misal Number, Complex, Ordered, dan lain-lain
--	https://pkg.go.dev/golang.org/x/exp/constraints

*/

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

/*
==	Maps & Slices Packages

--	Terdapat juga package maps dan slices, yang berisi function untuk mengelola data Map dan Slice, namun sudah menggunakan fitur Generic
--	https://pkg.go.dev/golang.org/x/exp/maps
--	https://pkg.go.dev/golang.org/x/exp/slices
*/
func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Mufied",
	}

	second := map[string]string{
		"Name": "Mufied",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlice(t *testing.T) {
	first := []string{"Ahmad"}
	second := []string{"Ahmad"}

	assert.True(t, slices.Equal(first, second))
}

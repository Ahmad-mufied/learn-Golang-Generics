package learn_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
==	Comparable

--	Selain any, di Go-Lang versi 1.18 juga terdapat tipe data bernama comparable
--	comparable merupakan interface yang diimplementasikan oleh tipe data yang bisa dibandingkan (menggunakan operator != dan ==), seperti booleans, numbers, strings, pointers, channels, interfaces, array yang isinya ada comparable type, atau structs yang fields nya adalah comparable type

*/

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Ahmad", "Ahmad"))
	assert.True(t, IsSame[int](100, 100))
	assert.True(t, IsSame[bool](true, true))
}

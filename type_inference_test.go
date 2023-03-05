package learn_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
==	Type Inference

--	Type Inference merupakan fitur dimana kita tidak perlu menyebutkan Type Parameter ketika memanggil kode generic
--	Tipe data Type Parameter bisa dibaca secara otomatis misal dari parameter yang kita kirim
--	Namun perlu diingat, pada beberapa kasus, jika terjadi error karena Type Inference, kita bisa dengan mudah memperbaikinya dengan cara menyebutkan Type Parameter nya saja

*/

func TestMinInference(t *testing.T) {
	assert.Equal(t, 100, Min1(100, 200))
	assert.Equal(t, int64(100), Min1(int64(100), int64(200)))
	assert.Equal(t, float64(100), Min1(float64(100), float64(200)))
	assert.Equal(t, Age(100), Min1(Age(100), Age(200)))
}

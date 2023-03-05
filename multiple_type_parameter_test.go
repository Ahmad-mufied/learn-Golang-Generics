package learn_golang_generics

import (
	"fmt"
	"testing"
)

/*
==	Multiple Type Parameter

--	Penggunakan Type Parameter bisa lebih dari satu, jika kita ingin menambahkan multiple Type Parameter, kita cukup gunakan tanda , (koma) sebagai pemisah antar Type Parameter
--	Nama Type Parameter harus berbeda, tidak boleh sama jika kita menambah Type Parameter lebih dari satu

*/

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Ahmad", 100)
	MultipleParameter[int, string](100, "Ahmad")
}

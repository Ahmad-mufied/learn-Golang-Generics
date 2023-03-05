package learn_golang_generics

import (
	"fmt"
	"testing"
)

/*
==	Generic Type

--	Sebelumnya kita sudah bahas tentang generic di function
--	Generic juga bisa digunakan ketika membuat type

*/

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"Ahmad", "Mufied", "Nugroho"}
	PrintBag[string](names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag[int](numbers)
}

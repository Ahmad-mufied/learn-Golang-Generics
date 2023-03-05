package learn_golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
==	Generic Struct

--	Struct juga mendukung generic
--	Dengan menggunakan generic, kita bisa membuat Field dengan tipe data yang sesuai dengan Type Parameter

*/

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[E]) ChangeFirst(first E) E {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Ahmad",
		Second: "Mufied",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Ahmad",
		Second: "Mufied",
	}

	assert.Equal(t, "Nugroho", data.ChangeFirst("Nugroho"))
	assert.Equal(t, "Hello Ahmad", data.SayHello("Ahmad"))
}

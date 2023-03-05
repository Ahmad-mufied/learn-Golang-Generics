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

/*
==	Generic Method

--	Selain di function, kita juga bisa tambahkan generic di method (function di struct)
--	Namun, generic di method merupakan generic yang terdapat di struct nya.
--	Kita wajib menyebutkan semua type parameter yang terdapat di struct, walaupun tidak kita gunakan misalnya, atau jika tidak ingin kita gunakan, kita bisa gunakan _ (garis bawah) sebagai pengganti type parameter nya
--	Method tidak bisa memiliki type parameter yang mirip dengan di function

*/

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Ahmad",
		Second: "Mufied",
	}

	assert.Equal(t, "Nugroho", data.ChangeFirst("Nugroho"))
	assert.Equal(t, "Hello Ahmad", data.SayHello("Ahmad"))
}

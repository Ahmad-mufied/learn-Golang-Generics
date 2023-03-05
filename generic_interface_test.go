package learn_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
==	Generic Interface

--	Generic juga bisa kita gunakan di Interface
--	Secara otomatis, semua struct yang ingin mengikuti kontrak interface tersebut harus menggunakan generic juga

*/

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{}

	result := ChangeValue[string](&myData, "Ahmad")

	assert.Equal(t, "Ahmad", result)
	assert.Equal(t, "Ahmad", myData.Value)
}

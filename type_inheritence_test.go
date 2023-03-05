package learn_golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
==	Type Parameter Inheritance

--	Go-Lang sendiri sebenarnya tidak memiliki pewarisan, namun seperti kita ketahui, jika kita membuat sebuah type yang sesuai dengan kontrak interface, maka dianggap sebagai implementasi interface tersebut
--	Type Parameter juga mendukung hal serupa, kita bisa gunakan constraint dengan menggunakan interface, maka secara otomatis semua interface yang compatible dengan type constraint tersebut bisa kita gunakan

*/

type Employee interface {
	GetName() string
}

func GetName[T Employee](parameter T) string {
	return parameter.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Ahmad", GetName[Manager](&MyManager{Name: "Ahmad"}))
	assert.Equal(t, "Ahmad", GetName[VicePresident](&MyVicePresident{Name: "Ahmad"}))
}

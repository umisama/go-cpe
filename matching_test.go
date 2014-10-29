package cpe

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckDisjoint(t *testing.T) {
	i1, err := NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	i2, err := NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	assert.Equal(t, false, CheckDisjoint(i1, i2))

	i1, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_95"]`)
	assert.Nil(t, err)
	i2, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	assert.Equal(t, true, CheckDisjoint(i1, i2))
	return
}

func TestCheckEqual(t *testing.T) {
	i1, err := NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	i2, err := NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	assert.Equal(t, true, CheckEqual(i1, i2))

	i1, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_200?"]`)
	assert.Nil(t, err)
	i2, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	assert.Equal(t, false, CheckEqual(i1, i2))
	return
}

func TestCheckSubset(t *testing.T) {
	i1, err := NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000",update="sp3"]`)
	assert.Nil(t, err)
	i2, err := NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	assert.Equal(t, true, CheckSubset(i1, i2))

	i1, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_95"]`)
	assert.Nil(t, err)
	i2, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	assert.Equal(t, false, CheckSubset(i1, i2))

	return
}

func TestCheckSuperset(t *testing.T) {
	i1, err := NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	i2, err := NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000",update="sp3",edition="pro"]`)
	assert.Nil(t, err)
	assert.Equal(t, true, CheckSuperset(i1, i2))

	i1, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_200*"]`)
	assert.Nil(t, err)
	i2, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000"]`)
	assert.Nil(t, err)
	assert.Equal(t, true, CheckSuperset(i1, i2))

	i1, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_95"]`)
	assert.Nil(t, err)
	i2, err = NewItemFromWfn(`wfn:[part="o",vendor="microsoft",product="windows_2000",update="sp3",edition="pro"]`)
	assert.Nil(t, err)
	assert.Equal(t, false, CheckSuperset(i1, i2))
}

func BenchmarkSuperset(b *testing.B) {
	item1, _ := NewItemFromUri("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")
	item2, _ := NewItemFromUri("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")
	b.ResetTimer()
	for i:=0; i< b.N; i++ {
		CheckSuperset(item1, item2)
	}
}

func BenchmarkSubset(b *testing.B) {
	item1, _ := NewItemFromUri("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")
	item2, _ := NewItemFromUri("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")
	b.ResetTimer()
	for i:=0; i< b.N; i++ {
		CheckSubset(item1, item2)
	}
}

func BenchmarkEqual(b *testing.B) {
	item1, _ := NewItemFromUri("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")
	item2, _ := NewItemFromUri("cpe:/a:microsoft:internet_explorer:8.0.6001:beta")
	b.ResetTimer()
	for i:=0; i< b.N; i++ {
		CheckEqual(item1, item2)
	}
}

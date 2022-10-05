package cases

import (
	"fmt"
)

type FiboStore interface {
	GetSequence(x int) (seq []int, err error)
	CheckParam(x int) error
}
type Fibo struct{
	fiboStore FiboStore
}

func NewFibo (fiboStore FiboStore) *Fibo{
	return &Fibo{
		fiboStore: fiboStore,
	}
}

func(f *Fibo) GetSequence(num int) ([]int, error){

	err := f.fiboStore.CheckParam(num)
	if err != nil {
		return nil, err
	}

	sl := []int{}
	a := 0
	b := 1
	c := b
	sl = append(sl, a, b)

	for true {
		c = b
		b = a + b
		if b >= num {
			break
		}
		a = c
		sl = append(sl, b)
	}
	return sl, nil
}



func CheckParam(num int) error{
	if num < 0 {
		return fmt.Errorf("Less 0")
	}
	return nil
}
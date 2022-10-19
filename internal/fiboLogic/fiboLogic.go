package fibologic

import (
	"errors"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/sirupsen/logrus"
)


type Fibo struct{
	Cache *memcache.Client
}

func(f *Fibo) GetSequence(x int, y int) ([]int64, error){

	err := f.CheckParam(x, y)
	if err != nil {
		return nil, err
	}
	sl := []int64{}

	for i := x; i <= y; i++{
		sl[i-x], err = f.FindNumber(i)
		if err != nil {
			logrus.Error("Cannot get fibonacci numbe")
		}
	}
	return sl, nil
}

func (f *Fibo)FindNumber(num int) (int64, error){
	//find fibonacci number in cache
	g, _ := f.Cache.Get(strconv.Itoa(num))
	if g != nil{
		return strconv.ParseInt(string(g.Value), 10, 64)
	}
	
	var r int64
	if num == 0 {
		r = 0
	}else if num < 2 {
		r = 1
	} else {
		var a int64
		var b int64
		a = 1
		b = 1
		
		for i := 2; i < num; i++{
			b = a + b 
			a = b - a
		}
		r = b
	}

	//write a number to cache 
	err := f.Cache.Set(&memcache.Item{Key: strconv.Itoa(num), Value: []byte(strconv.FormatInt(r, 10))})
	if err != nil {
		logrus.Warn("Cannot write to memcache", err)
	}
	return r, nil
}


func(f *Fibo) CheckParam(x int, y int) error{
	if x < 0 || y < 0 {
		return errors.New("Cannot be less than 0")
	}
	if x > y {
		return errors.New("Y cannot be leaa than x")
	}
	return nil
}
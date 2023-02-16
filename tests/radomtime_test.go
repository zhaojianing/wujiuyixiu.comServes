package test

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
	t.Log("A")
}
func TestRadomTimeNum(t *testing.T) {
	nowTime := time.Now().Unix()
	println('1', nowTime)
	t.Log('1', nowTime)
	nowTime1 := time.Now().Unix()
	t.Log('2', nowTime1)
}

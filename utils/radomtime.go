package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// RadomTimeNum 一个随机不重复的字符串
func RadomTimeNum() (str string, err error) {
	rand.NewSource(time.Now().Unix())
	now := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(9999))
	return now, nil
}

package util

import (
	"math/rand"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyz123456789-")

// RandString 随便获取指定长度的字符串
func RandString(n int) string {
	b := make([]rune, n)
	runesLen := len(runes)
	for i := range b {
		b[i] = runes[rand.Intn(runesLen)]
	}
	return string(b)
}

// Uniqid 生成唯一requestid
func Uniqid() string {
	return RandString(36)
}

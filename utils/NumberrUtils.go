package utils

import "math/rand"

func RandBetween(start int, end int) int {
RandNumber: // 随机暂停时间
	number := rand.Intn(end)
	if number < start {
		goto RandNumber
	}
	return number
}

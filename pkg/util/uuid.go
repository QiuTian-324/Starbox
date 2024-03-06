package util

import "github.com/google/uuid"

func GenUUID(l int) (uid string) {
	// 生成一个新的UUID
	newUUID := uuid.New()

	// 将UUID转换为10位长度的子串
	return newUUID.String()[:l]
}

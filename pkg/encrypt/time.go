package encrypt

import "time"

func FormatUpdateTime(updateTime time.Time) (updateTimeStr string) {
	// 使用 Format 方法将其转换为字符串
	return updateTime.Format("2006-01-02 15:04:05")
}

func FormatCreateTime(createTime time.Time) (updateTimeStr string) {
	return createTime.Format("2006-01-02 15:04:05")
}

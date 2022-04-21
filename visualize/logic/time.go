package logic

import (
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05"
)

// 处理时间相关函数
func TimeFormatString(t time.Time) string {
	timeStamp := t.Unix()
	fmt.Printf("%v\n\n\n", time.Unix(timeStamp, 0).Format(timeLayout))
	return time.Unix(timeStamp, 0).Format(timeLayout)
}

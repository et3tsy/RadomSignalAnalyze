package models

import "time"

// 定义模型信号
type Signal struct {
	Value      int64
	CreateTime time.Time
}

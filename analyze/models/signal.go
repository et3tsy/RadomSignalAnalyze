package models

import "time"

// 定义模型信号
type Signal struct {
	Value      int64     // 随机信号
	CreateTime time.Time // 信号产生时间
}

// 定义处理结果
type Result struct {
	Average    float64   // 一组信号的期望
	Variance   float64   // 一组信号的样本方差
	CreateTime time.Time // 该组信号中最晚产生的信号对应的时间
}

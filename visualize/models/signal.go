package models

import "time"

/*
	遇到的问题:
		消息队列中的内容都被取出了,但是读取时发现全是空值

	解决:
		models层中的模型均要定义两个,像Result和ParamResult

	原因:
		在rabbitMQ处理过程中,我们传递的是json格式的数据,在反射的时候,`json`的标识会让值映射json标志的属性
		进而我们真正应该接收变量就成了空值
*/

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

// 定义处理结果
type ParamResult struct {
	Average    float64 `json:"aver"`        // 一组信号的期望
	Variance   float64 `json:"vari"`        // 一组信号的样本方差
	CreateTime string  `json:"create_time"` // 该组信号中最晚产生的信号对应的时间
}

// 向前端传递统计模型
// 范围落在[Lowerbound,Upperbound)的次数Record
type ParamSignals struct {
	Lowerbound int64 `json:"lo"`
	Upperbound int64 `json:"hi"`
	Record     int64 `json:"value"`
}

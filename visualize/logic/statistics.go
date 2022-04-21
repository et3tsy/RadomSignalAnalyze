package logic

import (
	"visual/ds"
	mq "visual/messageQueue"
	"visual/models"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	t      *ds.Tree      // 定义线段树
	result models.Result // 定义统计结果
)

// 初始化线段树
func Init() {
	t = ds.NewTree(
		viper.GetInt64("signal.min"),
		viper.GetInt64("signal.max"),
	)
}

// 监听消息并且进行更新
func ReceiveSignalAndUpdate() {
	for {
		signal, err := mq.GetSignal()
		if err != nil {
			zap.L().Sugar().Error(err)
			continue
		}
		t.Update(signal.Value)
	}
}

// 将范围为[lo,hi]的数值进行划分,划分成均匀的段,段数为count
// 如果划分为[l1,r1),[l2,r2)...
// 将返回[l1,l2...]
func cutRange(lo, hi, count int64) (res []int64) {
	if count == 0 {
		return nil
	}

	// 设置切片大小
	res = make([]int64, count)

	// 分配大小
	eachSize := (hi - lo + 1) / count
	for i := range res {
		res[i] = eachSize
	}

	// 剩下多余的均匀分配
	lft := (hi - lo + 1) - count*eachSize
	for i := int64(0); i < lft; i++ {
		res[i]++
	}

	// 计算前缀和
	for i := int64(1); i < count; i++ {
		res[i] += res[i-1]
	}

	// 计算排名
	for i := count - 1; i > 0; i-- {
		res[i] = res[i-1] + lo
	}

	// 计数从lo开始
	res[0] = lo
	return
}

// 获得[lo,ri]的统计数据
func GetSignalStatistics(lo, hi, count int64) (ps []models.ParamSignals) {
	ps = make([]models.ParamSignals, count)
	res := cutRange(lo, hi, count)
	for i := int64(0); i < count-1; i++ {
		ps[i].Lowerbound = res[i]
		ps[i].Upperbound = res[i+1]
		ps[i].Record = t.Query(res[i], res[i+1]-1)
	}
	ps[len(ps)-1].Lowerbound = res[len(ps)-1]
	ps[len(ps)-1].Upperbound = hi + 1
	ps[len(ps)-1].Record = t.Query(res[len(ps)-1], hi)
	return
}

// 监听消息并且进行更新
func ReceiveResultAndUpdate() {
	for {
		var err error
		result, err = mq.GetResult()
		if err != nil {
			zap.L().Sugar().Error(err)
			continue
		}
	}
}

// 获得分析结果
func GetResult() *models.Result {
	return &result
}

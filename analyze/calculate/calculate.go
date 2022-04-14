package calculate

import (
	"analyze/queue"
	"fmt"
	"math"

	"github.com/spf13/viper"
)

var (
	size        int          // 需要分析最近信号的数量
	signalQueue *queue.Queue // 信号队列
	average     float64      // 平均数
	variance    float64      // 方差
)

// 初始化
func Init() {
	signalQueue = queue.NewQueue()
	size = viper.GetInt("analyze.size")
}

// 更新平均值
func calNewAverage(value int64, oldSize, newSize int) float64 {
	if newSize > oldSize {
		return (average*float64(oldSize) + float64(value)) / float64(newSize)
	} else {
		return (average*float64(oldSize) - float64(value)) / float64(newSize)
	}
}

// 更新方差,这里我们采用动态维护的方法,复杂度为O(1)
func calNewVariance(oldAverage float64, value int64, oldSize, newSize int) float64 {
	// 如果小于两个元素,方差为NaN
	if newSize < 2 {
		return 0
	}

	if newSize > oldSize {
		// 插入值

		// 刚好两个数,取出队首进行计算
		if newSize == 2 {
			v, _ := signalQueue.Front().(int64)
			return math.Pow(float64(v)-average, 2) + math.Pow(float64(value)-average, 2)
		}

		// 处理其他情况
		return (variance*float64(oldSize-1) +
			float64(oldSize)*math.Pow(oldAverage, 2) +
			math.Pow(float64(value), 2) -
			float64(newSize)*math.Pow(average, 2)) / float64(newSize-1)
	} else {
		return (variance*float64(oldSize-1) +
			float64(oldSize)*math.Pow(oldAverage, 2) -
			math.Pow(float64(value), 2) -
			float64(newSize)*math.Pow(average, 2)) / float64(newSize-1)
	}
}

// 插入新的元素
func Push(value int64) (err error) {
	signalQueue.Push(value)
	oldAverage := average
	average = calNewAverage(value, signalQueue.Size()-1, signalQueue.Size())
	variance = calNewVariance(oldAverage, value, signalQueue.Size()-1, signalQueue.Size())
	if signalQueue.Size() > size {
		iter := signalQueue.Pop()
		value, ok := iter.(int64)
		if !ok {
			return fmt.Errorf("[calculate]Fail to fetch the popped value")
		}
		oldAverage = average
		average = calNewAverage(value, signalQueue.Size()+1, signalQueue.Size())
		variance = calNewVariance(oldAverage, value, signalQueue.Size()+1, signalQueue.Size())
	}
	return nil
}

// 获取期望
func GetAverage() float64 {
	return average
}

// 获取方差
func GetVariance() float64 {
	if size <= 1 {
		return math.NaN()
	}
	return variance
}

package random

import (
	"math"
	"math/rand"
	"time"
)

// 正态分布
func GetGaussRandomNum(min, max int64) int64 {
	σ := (float64(min) + float64(max)) / 2
	μ := (float64(max) - σ) / 3
	rand.Seed(time.Now().UnixNano())
	x := rand.Float64()
	x1 := rand.Float64()
	a := math.Cos(2*math.Pi*x) * math.Sqrt((-2)*math.Log(x1))
	result := a*μ + σ
	return int64(result)
}

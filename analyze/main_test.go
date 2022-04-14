package main

import (
	"analyze/calculate"
	"log"
	"testing"
)

// 测试期望,方差计算
func TestPush(t *testing.T) {
	log.Println("TestPush")
	calculate.Push(5)
	calculate.Push(20)
	calculate.Push(40)
	calculate.Push(80)
	calculate.Push(100)
	want := 49.0
	got := calculate.GetAverage()
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}

	want = 1605
	got = calculate.GetVariance()
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}

	// 进行下面测试前先将confit.yaml中analyze.size设置为5
	calculate.Push(5)
	calculate.Push(20)
	calculate.Push(20)

	want = 49.0
	got = calculate.GetAverage()
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}

	want = 1605
	got = calculate.GetVariance()
	if got != want {
		t.Errorf("want:%v got:%v", want, got)
	}
}

func TestMain(m *testing.M) {
	log.Println("TestMain")
	setup()
	defer close()

	m.Run()
}

package logic

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestCutRange(t *testing.T) {
	log.Println("TestCutRange")
	r := cutRange(1, 10, 3)
	fmt.Printf("%v", r)
	r = cutRange(10, 20, 4)
	fmt.Printf("%v", r)
}

func TestTimeToString(t *testing.T) {
	timeStamp := time.Now().Unix()
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(timeStamp, 0).Format(timeLayout)
	fmt.Println(timeStr)
}

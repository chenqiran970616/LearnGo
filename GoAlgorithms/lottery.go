package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Reward struct {
	Key      string // 商品名称
	NumStart int    // 区间开始值
	NumEnd   int    // 区间结束值
}

// 权重抽奖逻辑 返回抽奖的key

func GetRandomWheel(r []Reward) string {
	if len(r) == 0 {
		panic("奖品不能为空!")
	}
	min := r[0].NumStart
	max := r[0].NumEnd
	key := ""
	// 获取最大最小值
	for index, item := range r {
		if index > 0 {
			if item.NumStart <= min {
				min = item.NumStart
			}
			if item.NumStart >= max {
				max = item.NumEnd
			}
		}
	}
	ts := time.Now().Unix()
	rand.Seed(ts)
	rt := rand.Intn(max-min) + min
	for _, item2 := range r {
		if rt >= item2.NumStart && rt <= item2.NumEnd {
			key = item2.Key
			fmt.Println("命中,", rt, item2)
			break
		}
	}
	//如果需要保持100%中奖概率，则必须保证抽奖key不能为空,而且必须有兜底奖品，否则会爆炸
	if key == "" {
		return GetRandomWheel(r)
	}

	return key
}
func main() {
	Rewards := []Reward{
		{Key: "mba15", NumStart: 1, NumEnd: 2},
		{Key: "mba16", NumStart: 3, NumEnd: 4},
		{Key: "mba17", NumStart: 5, NumEnd: 6},
		{Key: "很遗憾你没中奖", NumStart: 7, NumEnd: 100},
	}
	var BoeTeam = []string{1: "陈祺然", 2: "刘永为", 3: "郭全明", 4: "陕卓莹", 5: "孙振为", 6: "赵家兴", 7: "姚鑫琳", 8: "朱萌"}
	for i := 1; i < 9; i++ {
		fmt.Printf("第%d轮抽奖", i)
		fmt.Println(BoeTeam[i])
		GetRandomWheel(Rewards)
		time.Sleep(time.Duration(1) * time.Second)
	}
}

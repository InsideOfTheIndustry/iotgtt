package testrand

import (
	"fmt"
	"math/rand"
	"time"
)

// Generaterand 生成随机数
// 每次程序启动会产生与上一次程序启动时一样的随机数结果
func Generaterand() int {
	randnum1 := rand.Intn(100)
	randnum2 := rand.Intn(100)
	fmt.Println(randnum1)
	fmt.Println(randnum2)
	return randnum1
}

// Generaterandwithsameseed rand.seed相同时结果与上面一致
func Generaterandwithsameseed() int {
	rand.Seed(10)
	randnum1 := rand.Intn(100)
	randnum2 := rand.Intn(100)
	fmt.Println(randnum1)
	fmt.Println(randnum2)
	return randnum1
}

// Generaterandwithdiffseed rand.seed不同时生成的随机值会随着变化
func Generaterandwithdiffseed() int {
	rand.Seed(time.Now().UnixNano())
	randnum1 := rand.Intn(100)
	randnum2 := rand.Intn(100)
	fmt.Println(randnum1)
	fmt.Println(randnum2)
	return randnum1
}

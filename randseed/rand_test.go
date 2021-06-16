package testrand

import (
	"testing"
)

func TestGenerateRand(t *testing.T) {
	rand1 := GenerateRand()
	// 我的本地虚拟机中产生的随机数为 81
	if rand1 != 81 {
		t.Errorf("运行时应该会得到%v, 实际上得到:%v", 81, rand1)
	}
}

func TestGenerateRandWithSameSeed(t *testing.T) {
	rand1 := GenerateRandWithSameSeed()

	// 我的本地虚拟机中产生的随机数为 54
	if rand1 != 54 {
		t.Errorf("运行时应该会得到%v, 实际上得到:%v", 54, rand1)
	}
}

func TestGenerateRandWithDiffSeed(t *testing.T) {
	rand1 := GenerateRandWithDiffSeed()

	if rand1 == 85 {
		t.Errorf("本不应该生成随机数为%v的，但是还是生成了%v,当然也有可能是真的随机到了85", 85, rand1)
	}
}

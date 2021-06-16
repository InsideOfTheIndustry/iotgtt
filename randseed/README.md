# 测试rand

1、 若不使用 rand.seed 则在程序多次启动时，产生的随机数不会发生变化
2、 使用相同的 rand.seed 值， 产生的随机数也不会发生变化 如 rand.seed(10)
3、 使用不同的 rand.seed 值， 产生的随机数会发生变化 如 rand.seed(time.Now().UnixNano())
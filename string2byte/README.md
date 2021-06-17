#  []byte to string

## why

We know that string is a immutable type in Go language. Type conversation between string to []byte will trigger memory allocation. 

So we can point the original pointer to new datastruct.

我们知道string是一个不可改变的类型在Golang中，类型的转换会产生新的内存分配。

我们只需要修改底层数组指向新的数据类型即可。

## result 

```
======================
Convert Byte to String
======================
result:
goos: darwin
goarch: amd64
pkg: github.com/InsideOfTheIndustry/iotgtt/string2byte
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
BenchmarkOriginByte2Str-4      	219867090	         5.379 ns/op	       0 B/op	       0 allocs/op
BenchmarkConverStrToByte-4   	1000000000	         0.7359 ns/op	       0 B/op	       0 allocs/op
PASS
ok      github.com/InsideOfTheIndustry/iotgtt/string2byte       2.889s
```

```
======================
Convert String to []byte
======================
result:
goos: darwin
goarch: amd64
pkg: github.com/InsideOfTheIndustry/iotgtt/string2byte
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
BenchmarkOriginStr2Byte-4     	174552756	         6.785 ns/op	       0 B/op	       0 allocs/op
BenchmarkConverStr2Byte-4   	1000000000	         0.3185 ns/op	       0 B/op	       0 allocs/op
PASS
ok      github.com/InsideOfTheIndustry/iotgtt/string2byte       2.889s
```

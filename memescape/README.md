# Memory escape
Though the benchmark, we can see that return a pointer is not faster than return a value. 
In conventional thoughts as C language, pass Pointer is faster than pass a copy of struct value.

```
============
Benchmark 
============
goos: darwin
goarch: amd64
pkg: github.com/InsideOfTheIndustry/iotgtt/memescape
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
BenchmarkNewPersonWithPointer-4   	  25721979	          46.33 ns/op	      48 B/op	       1 allocs/op
BenchmarkNewPersonWithValue-4   	1000000000	         0.3313 ns/op	       0 B/op	       0 allocs/op
PASS
ok      github.com/InsideOfTheIndustry/iotgtt/memescape 2.414s
```

## why
Why? In Golang pass pointer may cause memory escape from stack to heap.


参照build文件夹下，进行逃逸分析。
```shell
go build -gcflags=-m 


# github.com/InsideOfTheIndustry/iotgtt/memescape/build
./main.go:18:6: can inline NewPersonWithPointer
./main.go:6:6: can inline main
./main.go:8:22: inlining call to NewPersonWithPointer
./main.go:8:22: new(Person) does not escape
./main.go:18:36: leaking param: name
./main.go:18:42: leaking param: hobby
./main.go:19:15: new(Person) escapes to heap
```

可以看到 `new(Person) escapes to heap`

传递指针可以减少底层值的拷贝，可以提高效率，但是如果拷贝的数据量小，
由于指针传递会产生逃逸，可能会使用堆，也可能会增加GC的负担，所以传递指针不一定是高效的

不要盲目使用变量的指针作为函数参数，虽然它会减少复制操作。
但其实当参数为变量自身的时候，复制是在栈上完成的操作，开销远比变量逃逸后动态地在堆上分配内存少的多

## more
- 对于小型的数据，使用传值而不是传指针，避免内存逃逸。
- 避免使用长度不固定的slice切片，在编译期无法确定切片长度，只能将切片使用堆分配。
- interface调用方法会发生内存逃逸，在热点代码片段，谨慎使用。

## Read!
https://eddycjy.gitbook.io/golang/di-1-ke-za-tan/stack-heap
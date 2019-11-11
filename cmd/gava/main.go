package main

import "jvm/pkg/gava"

func main() {
	// https://github.com/etcd-io/etcd/blob/b852e37895/main.go
	// 为方便单元测试，参考 etcd 实现
	// 单元测试见 test/pkg/classpath/classpath_test.go
	gava.Main()
}

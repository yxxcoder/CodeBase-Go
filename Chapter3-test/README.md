# Chapter3-test 单元测试与性能测试
单元测试与性能测试的方法，详情请见calculate.go,calculate_test.go
### 1. 执行单元测试
```shell
go test -v
```
### 2. 执行性能测试
go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench
```shell
# 测试全部的压力测试函数
go test -test.bench=".*"
```
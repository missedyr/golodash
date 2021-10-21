# golodash

##
###  使用方法
go get https://github.com/xuexin520/golodash

##
### 数组 切片 集合
* [func  IndexOf](#funcIndexOf)

##
### 字符串


###  <a name='funcIndexOf'></a>func  IndexOf
```go
func IndexOf(items DashSlice, element interface{}) (int, bool)
```
This method is like _.find except that it returns the index of the first element
predicate returns truthy for instead of the element itself.
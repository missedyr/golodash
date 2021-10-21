# golodash

##
###  使用方法
go get https://github.com/xuexin520/golodash

##
### 数组 切片 集合
* [func  IndexOfInt](#funcIndexOfInt)
* [func  IndexOfInt8](#funcIndexOfInt8)
* [func  IndexOfInt32](#funcIndexOfInt32)
* [func  IndexOfInt64](#funcIndexOfInt64)
* [func  IndexOfInterface](#funcIndexOfInterface)

##
### 字符串


###  <a name='funcIndexOfInt'></a>func  IndexOfInt
```
func IndexOfInt(items DashSliceInt, value int) int

items:  an slice of values of type int
Return: the first index value found in array
```
#### Returns the first index value found in array array

###  <a name='funcIndexOfInt8'></a>func  IndexOfInt8
```
func IndexOfInt8(items DashSliceInt8, value int8) int
```

###  <a name='funcIndexOfInt32'></a>func  IndexOfInt32
```
func IndexOfInt32(items DashSliceInt32, value int32) int
```

###  <a name='funcIndexOfInt64'></a>func  IndexOfInt64
```
func IndexOfInt64(items DashSliceInt64, value int64) int
```

###  <a name='funcIndexOfInterface'></a>func  IndexOfInterface
```
func IndexOfInterface(items DashSliceInterface, value interface{}) int
```

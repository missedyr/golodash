# golodash

##
###  使用方法
go get github.com/missedyr/golodash

##
### 数组 切片 集合
* [func  IndexOfString](#funcIndexOfString)
* [func  IndexOfInt](#funcIndexOfInt)
* [func  IndexOfInt8](#funcIndexOfInt8)
* [func  IndexOfInt32](#funcIndexOfInt32)
* [func  IndexOfInt64](#funcIndexOfInt64)
* [func  IndexOfInterface](#funcIndexOfInterface)
* [func  ConcatOfInt](#funcConcatOfInt)
* [func  ConcatOfString](#funcConcatOfString)
* [func  ConcatOfInterface](#funcConcatOfInterface)
##

### 字符串

##

### 对象 map
* [func  Assign](#funcAssign)

##

### func 描述
###  <a name='funcIndexOfString'></a> func IndexOfString
```
func IndexOfString(items DashSliceString, value string) int

Items:  an slice of values of type string
Return(int): the first index value found in array, No found for return -1
```

###  <a name='funcIndexOfInt'></a> func IndexOfInt
```
func IndexOfInt(items DashSliceInt, value int) int

Items:  an slice of values of type int
Return(int): the first index value found in array, No found for return -1
```

###  <a name='funcIndexOfInt8'></a> func IndexOfInt8
```
func IndexOfInt8(items DashSliceInt8, value int8) int

Items:  an slice of values of type int8
Return(int): the first index value found in array, No found for return -1
```

###  <a name='funcIndexOfInt32'></a> func  IndexOfInt32
```
func IndexOfInt32(items DashSliceInt32, value int32) int

Items:  an slice of values of type int32
Return(int): the first index value found in array, No found for return -1
```

###  <a name='funcIndexOfInt64'></a> func  IndexOfInt64
```
func IndexOfInt64(items DashSliceInt64, value int64) int

Items:  an slice of values of type int64
Return(int): the first index value found in array, No found for return -1
```

###  <a name='funcIndexOfInterface'></a> func IndexOfInterface  
```
func IndexOfInterface(items DashSliceInterface, value interface{}) int

Items:  an slice of values of type interface
Return(int): the first index value found in array, No found for return -1
```

****

###  <a name='funcConcatOfInt'></a> func ConcatOfInt
```
func ConcatOfInt(arrays ...[]int) []int

arrays:  an slice of values of type int
Return:  the new array after joining
```

****

###  <a name='funcConcatOfString'></a> func ConcatOfString
```
func ConcatOfString(arrays ...[]string) []string

arrays:  an slice of values of type string
Return:  the new array after joining
```

****

###  <a name='funcConcatOfInterface'></a> func ConcatOfInterface
```
func ConcatOfInterface(arrays ...[]interface{}) []interface{}

arrays:  an slice of values of type interface
Return:  the new array after joining
```

****

###  <a name='funcAssign'></a> func Assign
```
func Assign(mapObj interface{}, sources ...interface{}) map[string]interface{} 

mapObj (map): 目标对象。   map或struct值 不接收指针类型
[sources] (...map): 来源对象   map或struct值 不接收指针类型
Return(int): map[string]interface{}
```
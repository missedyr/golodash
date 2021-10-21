package lodash

func IndexOfInt(items DashSliceInt, value int) int {
	var index = -1
	for i, v := range items {
		if v == value {
			index = i
			break
		}
	}
	return index
}

func IndexOfInt8(items DashSliceInt8, value int8) int {
	var index = -1
	for i, v := range items {
		if v == value {
			index = i
			break
		}
	}
	return index
}

func IndexOfInt32(items DashSliceInt32, value int32) int {
	var index = -1
	for i, v := range items {
		if v == value {
			index = i
			break
		}
	}
	return index
}

func IndexOfInt64(items DashSliceInt64, value int64) int {
	var index = -1
	for i, v := range items {
		if v == value {
			index = i
			break
		}
	}
	return index
}

func IndexOfInterface(items DashSliceInterface, value interface{}) int {
	var index = -1
	for i, v := range items {
		if v == value {
			index = i
			break
		}
	}
	return index
}
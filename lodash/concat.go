package lodash

func ConcatOfInt(arrays ...[]int) []int {
	list := []int{}
	for _, arr := range arrays {
		for _, v := range arr {
			list = append(list, v)
		}
	}
	return list
}

func ConcatOfString(arrays ...[]string) []string {
	list := []string{}
	for _, arr := range arrays {
		for _, v := range arr {
			list = append(list, v)
		}
	}
	return list
}

func ConcatOfInterface(arrays ...[]interface{}) []interface{} {
	list := []interface{}{}
	for _, arr := range arrays {
		for _, v := range arr {
			list = append(list, v)
		}
	}
	return list
}
